package github

import (
	"fmt"

	"github.com/google/go-github/github"
)

// Service to parsing and formater from response
func ParseResponse(notif []*github.Notification) {
	for _, notifinfo := range notif {
		if *notifinfo.Reason == "security_alert" {
			fmt.Println("[+] Updated At	: ", *notifinfo.UpdatedAt)
			fmt.Println("[+] Repo 	: ", *notifinfo.Repository.HTMLURL)
			fmt.Println("[+] Owner	: ", *notifinfo.Repository.Owner.Login)
			fmt.Println("[+] Reason 	: ", *notifinfo.Reason)
			fmt.Println("[+] Title 	: ", *notifinfo.Subject.Title, "\n")
		}
	}
}
