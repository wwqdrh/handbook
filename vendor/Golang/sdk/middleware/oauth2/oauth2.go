package oauth2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Setup return an oauth2Config configured to talk
// to github, you need environment variables set
// for your id and secret
func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}

// GetToken retrieves a github oauth2 token
func GetToken(ctx context.Context, conf *oauth2.Config) (*oauth2.Token, error) {
	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}

// GetUsers uses an initialized oauth2 client to get
// information about a user
func GetUsers(client *http.Client) error {
	url := fmt.Sprintf("https://api.github.com/user")

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Status Code from", url, ":", resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)
	return nil
}

////////////////////
// token - storage
////////////////////

// Storage is our generic storage interface
type Storage interface {
	GetToken() (*oauth2.Token, error)
	SetToken(*oauth2.Token) error
}

// Config wraps the default oauth2.Config
// and adds our storage
type Config struct {
	*oauth2.Config
	Storage
}

// Exchange stores a token after retrieval
func (c *Config) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	if err := c.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// TokenSource can be passed a token which
// is stored, or when a new one is retrieved,
// that's stored
func (c *Config) TokenSource(ctx context.Context, t *oauth2.Token) oauth2.TokenSource {
	return StorageTokenSource(ctx, c, t)
}

// Client is attached to our TokenSource
func (c *Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return oauth2.NewClient(ctx, c.TokenSource(ctx, t))
}

type storageTokenSource struct {
	*Config
	oauth2.TokenSource
}

// GetToken retrieves a github oauth2 token
func GetTokenByConf(ctx context.Context, conf Config) (*oauth2.Token, error) {
	token, err := conf.Storage.GetToken()
	if err == nil && token.Valid() {
		return token, err
	}
	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil, err
	}
	return conf.Exchange(ctx, code)
}

// Token satisfies the TokenSource interface
func (s *storageTokenSource) Token() (*oauth2.Token, error) {
	if token, err := s.Config.Storage.GetToken(); err == nil && token.Valid() {
		return token, err
	}
	token, err := s.TokenSource.Token()
	if err != nil {
		return token, err
	}
	if err := s.Config.Storage.SetToken(token); err != nil {
		return nil, err
	}
	return token, nil
}

// StorageTokenSource will be used by out configs TokenSource
// function
func StorageTokenSource(ctx context.Context, c *Config, t *oauth2.Token) oauth2.TokenSource {
	if t == nil || !t.Valid() {
		if tok, err := c.Storage.GetToken(); err == nil {
			t = tok
		}
	}
	ts := c.Config.TokenSource(ctx, t)
	return &storageTokenSource{c, ts}
}

////////////////////
// file-storage
////////////////////

// FileStorage satisfies our storage interface
type FileStorage struct {
	Path string
	mu   sync.RWMutex
}

// GetToken retrieves a token from a file
func (f *FileStorage) GetToken() (*oauth2.Token, error) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	in, err := os.Open(f.Path)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	var t *oauth2.Token
	data := json.NewDecoder(in)
	return t, data.Decode(&t)
}

// SetToken creates, truncates, then stores a token
// in a file
func (f *FileStorage) SetToken(t *oauth2.Token) error {
	if t == nil || !t.Valid() {
		return errors.New("bad token")
	}

	f.mu.Lock()
	defer f.mu.Unlock()
	out, err := os.OpenFile(f.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()
	data, err := json.Marshal(&t)
	if err != nil {
		return err
	}

	_, err = out.Write(data)
	return err
}
