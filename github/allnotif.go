package github

import (
	"context"
	"net/http"

	"fmt"

	"github.com/google/go-github/github"
) // with go modules disabled

// get notification in all repo
func GetNotifications(session *http.Client, csv bool) {
	ctx := context.Background()
	client := github.NewClient(session)

	// Get 1 month -> since
	since := TimeParse()
	// optional add for seach query in github
	opts := &github.NotificationListOptions{
		All:   true,
		Since: since,
		ListOptions: github.ListOptions{
			PerPage: 1000000,
		},
	}

	notif, _, err := client.Activity.ListNotifications(ctx, opts)
	if err != nil {
		fmt.Println(err)
	}

	ParseResponseNotif(notif)
	if csv {
		GenerateCsv(notif)
	}
}
