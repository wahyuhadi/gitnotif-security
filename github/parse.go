package github

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

// Service to parsing and formater from response
func ParseResponseNotif(notif []*github.Notification) {
	for _, notifinfo := range notif {
		if *notifinfo.Reason == "security_alert" {
			fmt.Println("[FOUND] Updated -> ", *notifinfo.UpdatedAt)
			fmt.Println("[FOUND] Repo 	-> ", *notifinfo.Repository.HTMLURL)
			fmt.Println("[FOUND] Owner	-> ", *notifinfo.Repository.Owner.Login)
			fmt.Println("[FOUND] Reason  -> ", *notifinfo.Reason)
			fmt.Println("[FOUND] Title 	-> ", *notifinfo.Subject.Title, "\n")
		}
	}
}

func ParseResponseListRepo(session *http.Client, repos []*github.Repository) {
	for _, repo := range repos {
		fmt.Println("[+] Repo Name Scan ", *repo.Name)
		NotifInRepo(session, *repo.Name)
	}
}

// Generate CSV  from response
func GenerateCsv(notif []*github.Notification) {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	strWrite := [][]string{{"Repo", "Owner", " Title", "Reason"}}
	for _, notifinfo := range notif {
		if *notifinfo.Reason == "security_alert" {
			add := []string{*notifinfo.Repository.HTMLURL, *notifinfo.Repository.Owner.Login, *notifinfo.Subject.Title, *notifinfo.Reason}
			strWrite = append(strWrite, add)
		}
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()

}
