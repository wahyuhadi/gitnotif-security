package github

import (
	"fmt"
	"os"
)

// get notif in Repo
func NotifInRepo(repo string) {
	owner := os.Getenv("owner")
	fmt.Println(owner, repo)
}
