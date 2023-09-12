package shell

import (
	"fmt"
	"strings"
)

type Middleware interface {
	Execute(cmd string) (string, string, error)
	Exit()
}

type session struct {
	upstream Middleware
	name     string
}

const (
	HTTPPort  = 5985
	HTTPSPort = 5986
)

type SessionConfig struct {
	ComputerName          string
	AllowRedirection      bool
	Authentication        string
	CertificateThumbprint string
	Credential            interface{}
	Port                  int
	UseSSL                bool
}


// utf8 implements a primitive middleware that encodes all outputs
// as base64 to prevent encoding issues between remote PowerShell
// shells and the receiver. Just setting $OutputEncoding does not
// work reliably enough, sadly.
type utf8 struct {
	upstream Middleware
	wrapper  string
}

func NewUTF8(upstream Middleware) (Middleware, error) {
	wrapper := "goUTF8" + utils.CreateRandomString(8)

	_, _, err := upstream.Execute(fmt.Sprintf(`function %s { process { if ($_) { [Convert]::ToBase64String([Text.Encoding]::UTF8.GetBytes($_)) } else { '' } } }`, wrapper))

	return &utf8{upstream, wrapper}, err
}

func (u *utf8) Execute(cmd string) (string, string, error) {
	// Out-String to concat all lines into a single line,
	// Write-Host to prevent line breaks at the "window width"
	cmd = fmt.Sprintf(`%s | Out-String | %s | Write-Host`, cmd, u.wrapper)

	stdout, stderr, err := u.upstream.Execute(cmd)
	if err != nil {
		return stdout, stderr, err
	}

	decoded, err := base64.StdEncoding.DecodeString(stdout)
	if err != nil {
		return stdout, stderr, err
	}

	return string(decoded), stderr, nil
}

func (u *utf8) Exit() {
	u.upstream.Exit()
}

func NewSessionConfig() *SessionConfig {
	return &SessionConfig{}
}

func (c *SessionConfig) ToArgs() []string {
	args := make([]string, 0)

	if c.ComputerName != "" {
		args = append(args, "-ComputerName")
		args = append(args, QuoteArg(c.ComputerName))
	}

	if c.AllowRedirection {
		args = append(args, "-AllowRedirection")
	}

	if c.Authentication != "" {
		args = append(args, "-Authentication")
		args = append(args, QuoteArg(c.Authentication))
	}

	if c.CertificateThumbprint != "" {
		args = append(args, "-CertificateThumbprint")
		args = append(args, QuoteArg(c.CertificateThumbprint))
	}

	if c.Port > 0 {
		args = append(args, "-Port")
		args = append(args, strconv.Itoa(c.Port))
	}

	if asserted, ok := c.Credential.(string); ok {
		args = append(args, "-Credential")
		args = append(args, asserted) // do not quote, as it contains a variable name when using password auth
	}

	if c.UseSSL {
		args = append(args, "-UseSSL")
	}

	return args
}

type credential interface {
	prepare(Middleware) (interface{}, error)
}

type UserPasswordCredential struct {
	Username string
	Password string
}

func (c *UserPasswordCredential) prepare(s Middleware) (interface{}, error) {
	name := "goCred" + CreateRandomString(8)
	pwname := "goPass" + CreateRandomString(8)

	_, _, err := s.Execute(fmt.Sprintf("$%s = ConvertTo-SecureString -String %s -AsPlainText -Force", pwname, QuoteArg(c.Password)))
	if err != nil {
		return nil, fmt.Errorf("%w, Could not convert password to secure string", err) errors.Annotate(err, "Could not convert password to secure string")
	}

	_, _, err = s.Execute(fmt.Sprintf("$%s = New-Object -TypeName 'System.Management.Automation.PSCredential' -ArgumentList %s, $%s", name, QuoteArg(c.Username), pwname))
	if err != nil {
		return nil, fmt.Errorf("%w, Could not create PSCredential object", err) errors.Annotate(err, "Could not create PSCredential object")
	}

	return fmt.Sprintf("$%s", name), nil
}

func NewSession(upstream Middleware, config *SessionConfig) (Middleware, error) {
	asserted, ok := config.Credential.(credential)
	if ok {
		credentialParamValue, err := asserted.prepare(upstream)
		if err != nil {
			return nil, fmt.Errorf("%w, Could not setup credentials", err)
		}

		config.Credential = credentialParamValue
	}

	name := "goSess" + CreateRandomString(8)
	args := strings.Join(config.ToArgs(), " ")

	_, _, err := upstream.Execute(fmt.Sprintf("$%s = New-PSSession %s", name, args))
	if err != nil {
		return nil, fmt.Errorf("%w, Could not create new PSSession")
	}

	return &session{upstream, name}, nil
}

func (s *session) Execute(cmd string) (string, string, error) {
	return s.upstream.Execute(fmt.Sprintf("Invoke-Command -Session $%s -Script {%s}", s.name, cmd))
}

func (s *session) Exit() {
	s.upstream.Execute(fmt.Sprintf("Disconnect-PSSession -Session $%s", s.name))
	s.upstream.Exit()
}
