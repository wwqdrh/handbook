package shell

import (
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

type (
	Waiter interface {
		Wait() error
	}

	Starter interface {
		StartProcess(cmd string, args ...string) (Waiter, io.Writer, io.Reader, io.Reader, error)
	}

	// sshSession exists so we don't create a hard dependency on crypto/ssh.
	sshSession interface {
		Waiter

		StdinPipe() (io.WriteCloser, error)
		StdoutPipe() (io.Reader, error)
		StderrPipe() (io.Reader, error)
		Start(string) error
	}
)

type Local struct{}

type SSH struct {
	Session sshSession
}

func (b *Local) StartProcess(cmd string, args ...string) (Waiter, io.Writer, io.Reader, io.Reader, error) {
	command := exec.Command(cmd, args...)

	stdin, err := command.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not get hold of the PowerShell's stdin stream", err)
	}

	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not get hold of the PowerShell's stdout stream")
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not get hold of the PowerShell's stderr stream")
	}

	err = command.Start()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not spawn PowerShell process")
	}

	return command, stdin, stdout, stderr, nil
}

func (b *SSH) StartProcess(cmd string, args ...string) (Waiter, io.Writer, io.Reader, io.Reader, error) {
	stdin, err := b.Session.StdinPipe()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not get hold of the SSH session's stdin stream", err)
	}

	stdout, err := b.Session.StdoutPipe()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not get hold of the SSH session's stdout stream", err)
	}

	stderr, err := b.Session.StderrPipe()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not get hold of the SSH session's stderr stream", err)
	}

	err = b.Session.Start(b.createCmd(cmd, args))
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("%w, Could not spawn process via SSH", err)
	}

	return b.Session, stdin, stdout, stderr, nil
}

func (b *SSH) createCmd(cmd string, args []string) string {
	parts := []string{cmd}
	simple := regexp.MustCompile(`^[a-z0-9_/.~+-]+$`)

	for _, arg := range args {
		if !simple.MatchString(arg) {
			arg = b.quote(arg)
		}

		parts = append(parts, arg)
	}

	return strings.Join(parts, " ")
}

func (b *SSH) quote(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}
