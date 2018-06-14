package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

const (
	appName    = "ghutils"
	appVersion = "0.2.0"
)

var (
	ctx    = context.Background()
	client *github.Client
)

func main() {

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "prints only the version",
	}

	app := cli.NewApp()
	app.Name = appName
	app.Usage = "GitHub helper utility"
	app.Compiled = time.Now()
	app.Version = appVersion
	app.HideHelp = false
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "GitHub Access Token",
			EnvVar: "GITHUB_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:  "org, o",
			Usage: "GitHub organization name",
		},
		cli.StringFlag{
			Name:  "user, u",
			Usage: "Github username",
		},
	}
	app.Before = configClient
	app.Action = defaultAction
	app.Commands = []cli.Command{
		addUserToTeamCommand,
		getNotificationsCommand,
		listOrgMemberCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\nError: %v\n\n", err)
	} else {
		fmt.Println()
	}

}

func configClient(c *cli.Context) error {
	token := c.GlobalString("token")
	if token == "" {
		return fmt.Errorf("access token required")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client = github.NewClient(oauth2.NewClient(ctx, ts))
	return nil
}

func defaultAction(c *cli.Context) error {
	org := c.String("org")
	username := c.String("user")

	if org != "" {
		return printTeams(org)
	}
	if username != "" {
		return printUser(username)
	}

	fmt.Println("No arguments provided, here are some samples...")

	fmt.Println("List teams in organization:")
	fmt.Printf("%s -o my-org-name\n\n", appName)

	fmt.Println("Print user details:")
	fmt.Printf("%s -u someuser\n\n", appName)

	fmt.Println("Add user to team:")
	fmt.Printf("%s add -user someuser -team 1234567\n\n", appName)

	fmt.Println("List notifications:")
	fmt.Printf("%s notif\n\n", appName)

	return nil
}
