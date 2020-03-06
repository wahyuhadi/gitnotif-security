package main

import (
	"context"
	"fmt"
	"gitnotif/github"
	"os"

	"golang.org/x/oauth2"
)

func main() {
	fmt.Println("[+] git notif security")
	token := os.Getenv("git_token")

	if token == "" {
		fmt.Println("[!]  No token found in .env { git_token }")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	github.GetNotifications(tc)
}
