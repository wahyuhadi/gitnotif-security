package github	

import (
	"context"
	"net/http"

	"fmt"
	"time"
	"github.com/google/go-github/github"
) // with go modules disabled

func GetNotifications(session *http.Client) {
	ctx := context.Background()
	client := github.NewClient(session)

	// Get 1 month -> since 
	now := time.Now()
	then := now.AddDate(0, 0, -1).Format("2006-01-02T15:04:05:0700")
	fmt.Println(then)
	since, _ := time.Parse("2006-01-02", then)
	
	fmt.Println(since)
	// optional add for seach quety in github
	opts := &github.NotificationListOptions{
		All: true,
		Since: since,
		ListOptions: github.ListOptions{
			PerPage : 1000000,
		},
	}

	notif , _, err := client.Activity.ListNotifications(ctx, opts)
	if err != nil {
		fmt.Println(err)
	}

	for _, notifinfo := range notif {
		if *notifinfo.Reason == "security_alert" {
			fmt.Println("[+] Updated At	: ", *notifinfo.UpdatedAt)
			fmt.Println("[+] Repo 	: ",*notifinfo.Repository.HTMLURL)
			fmt.Println("[+] Reason 	: ",*notifinfo.Reason)
			fmt.Println("[+] Title 	: ", *notifinfo.Subject.Title,"\n")
		}
	}
}