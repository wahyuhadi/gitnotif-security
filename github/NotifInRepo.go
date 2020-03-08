package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

// get notif in Repo
func NotifInRepo(session *http.Client, repo string) {
	owner := os.Getenv("owner")
	ctx := context.Background()
	client := github.NewClient(session)

	since := TimeParse()

	opts := &github.NotificationListOptions{
		All:   true,
		Since: since,
		ListOptions: github.ListOptions{
			PerPage: 1000000,
		},
	}

	notif, _, err := client.Activity.ListRepositoryNotifications(ctx, owner, repo, opts)
	if err != nil {
		fmt.Println(err)
	}

	ParseResponseNotif(notif)
}
