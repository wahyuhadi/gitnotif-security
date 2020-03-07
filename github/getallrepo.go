package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

func RepoVerbose(session *http.Client) {

	ctx := context.Background()
	client := github.NewClient(session)

	// optional add for seach query in github
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			PerPage: 1000000,
		},
	}

	owner := os.Getenv("owner")
	repos, _, err := client.Repositories.ListByOrg(ctx, owner, opts)
	if err != nil {
		fmt.Println(err)
	}

	ParseResponseListRepo(session, repos)
}
