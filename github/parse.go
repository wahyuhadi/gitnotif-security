package github

import (
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

// Service to parsing and formater from response
func ParseResponseNotif(notif []*github.Notification) {
	for _, notifinfo := range notif {
		if *notifinfo.Reason == "security_alert" {
			fmt.Println("[FOUND] Updated At	: ", *notifinfo.UpdatedAt)
			fmt.Println("[FOUND] Repo 	: ", *notifinfo.Repository.HTMLURL)
			fmt.Println("[FOUND] Owner	: ", *notifinfo.Repository.Owner.Login)
			fmt.Println("[FOUND] Reason 	: ", *notifinfo.Reason)
			fmt.Println("[FOUND] Title 	: ", *notifinfo.Subject.Title, "\n")
		}
	}
}

func ParseResponseListRepo(session *http.Client, repos []*github.Repository) {
	for _, repo := range repos {
		fmt.Println("[+] Repo Name Scan ", *repo.Name)
		NotifInRepo(session, *repo.Name)
	}
}
