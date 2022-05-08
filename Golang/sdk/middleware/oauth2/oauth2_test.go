package oauth2

import (
	"context"
	"io"
	"os"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func TestGithubOauth2(t *testing.T) {
	ctx := context.Background()
	conf := Setup()

	tok, err := GetToken(ctx, conf)
	if err != nil {
		panic(err)
	}
	client := conf.Client(ctx, tok)

	if err := GetUsers(client); err != nil {
		panic(err)
	}
}

func TestGithubTokenStorage(t *testing.T) {
	conf := Config{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GITHUB_CLIENT"),
			ClientSecret: os.Getenv("GITHUB_SECRET"),
			Scopes:       []string{"repo", "user"},
			Endpoint:     github.Endpoint,
		},
		Storage: &FileStorage{Path: "token.txt"},
	}
	ctx := context.Background()
	token, err := GetTokenByConf(ctx, conf)
	if err != nil {
		panic(err)
	}

	cli := conf.Client(ctx, token)
	resp, err := cli.Get("https://api.github.com/user")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}
