package main

import (
	"flag"
	"fmt"
	gh "github.com/google/go-github/v32/github"
	"github.com/marcusirgens/github-notifications/github"
	"github.com/marcusirgens/github-notifications/httpclient"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"time"
)

var token string

func main() {
	t, err := getGithubToken()
	if err != nil {
		fmt.Println("Please set a GitHub API token either using -token or GITHUB_NOTIFICATIONS_TOKEN")
		os.Exit(1)
	}

	client := getAuthenticatedGithubClient(t)


	nots, err := github.GetRecentNotifications(client)
	if err != nil {
		fmt.Printf("Could not get notifications: %v\n", err)
		os.Exit(1)
	}

	printNotifications(nots)
}

// getAuthenticatedGithubClient returns an authenticated GitHub client for the
// provided API token
func getAuthenticatedGithubClient(t string) *gh.Client {
	return gh.NewClient(getAuthenticatedClient(t))
}

// get a oauth2 client for the provided token
func getAuthenticatedClient(token string) *http.Client {
	// get a token source
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})

	return httpclient.NewCachedOauthClient(time.Minute * 5, ts)
}

func getGithubToken () (string, error) {
	if token != "" {
		return token, nil
	}
	env := os.Getenv("GITHUB_NOTIFICATIONS_TOKEN")
	if env != "" {
		return env, nil
	}

	return "", fmt.Errorf("no token set")
}

func init() {
	flag.StringVar(&token, "token", "", "Set a GitHub API token")
	flag.Parse()
}