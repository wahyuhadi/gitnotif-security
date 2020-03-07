package main

import (
	"context"
	"flag"
	"fmt"
	"gitnotif/github"
	"os"

	"golang.org/x/oauth2"
)

var (
	repo    = flag.String("repo", "false", "for search in repo") // default value is false
	verbose = flag.Bool("verbose", false, "Verbose mode")
)

func main() {
	flag.Parse()

	token := os.Getenv("git_token")
	if token == "" {
		fmt.Println("[!]  No token found in .env { git_token }")
		return
	}

	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	if *verbose {
		github.RepoVerbose(tc)
	} else {
		if *repo == "false" {
			github.GetNotifications(tc)
		}

		if *repo != "false" {
			github.NotifInRepo(tc, *repo)
		}

	}
}
