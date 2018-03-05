package main

import (
	"context"
	"flag"
	"log"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

var (
	ctx    = context.Background()
	logger = log.New(os.Stdout, "[github] ", log.Lshortfile)
)

func main() {

	// CONFIG
	orgName := flag.String("org", "", "GitHub org name")
	flag.Parse()
	if *orgName == "" {
		logger.Fatalf("Missing required config: org:%v", *orgName)
	}
	// END CONFIG

	// AUTH
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")})
	client := github.NewClient(oauth2.NewClient(ctx, ts))
	// END AUTH

	// repos
	PrintRepos(client, *orgName)

	// teams and members
	PrintTeams(client, *orgName)
}
