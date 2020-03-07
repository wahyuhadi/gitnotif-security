package main

import (
	"context"
	"flag"
	"fmt"
	"gitnotif/github"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	repo    = flag.String("repo", "false", "for search in repo") // default value is false
	verbose = flag.Bool("verbose", false, "Verbose mode")

	web = flag.Bool("web", false, "web mode")
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

	if *web {
		WebServer()
	}

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

func WebServer() {
	router := gin.Default()

	//new template engine
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}
