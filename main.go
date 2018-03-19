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
	app.Before = func(c *cli.Context) error {
		return configure(c)
	}
	app.Action = func(c *cli.Context) error {
		org := c.String("org")
		username := c.String("user")

		if org != "" {
			return PrintTeams(org)
		}
		if username != "" {
			return PrintUser(username)
		}

		fmt.Println("No arguments provided, here are some samples...")

		fmt.Println("List teams in organization:")
		fmt.Printf("%s -o my-org-name\n\n", appName)

		fmt.Println("Print user details:")
		fmt.Printf("%s -u someuser\n\n", appName)

		fmt.Println("Add user to team:")
		fmt.Printf("%s add -user someuser -team 1234567\n\n", appName)

		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add user to GitHub team",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "team, t",
					Usage: "GitHub Team ID",
				},
				cli.StringFlag{
					Name:  "user, u",
					Usage: "GitHub username",
				},
			},
			Action: func(c *cli.Context) error {

				team := c.Int64("team")
				user := c.String("user")

				if team == 0 {
					return fmt.Errorf("team argument required")
				}

				if user == "" {
					return fmt.Errorf("user argument required")
				}

				return AddUserToTeam(team, user)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\nError: %v\n\n", err)
	} else {
		fmt.Println()
	}

}

func configure(c *cli.Context) error {
	token := c.GlobalString("token")
	if token == "" {
		return fmt.Errorf("access token required")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client = github.NewClient(oauth2.NewClient(ctx, ts))
	return nil
}
