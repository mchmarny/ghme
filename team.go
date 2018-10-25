package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

var (
	addUserToTeamCommand = cli.Command{
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

			return addUserToTeam(team, user)
		},
	}
)

func printTeams(org string) error {

	if org == "" {
		return fmt.Errorf("org argument required")
	}

	fmt.Println()

	opt := &github.ListOptions{PerPage: 10}
	var allItems []*github.Team
	for {
		list, resp, err := client.Teams.ListTeams(ctx, org, opt)
		if err != nil {
			return err
		}
		allItems = append(allItems, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, e := range allItems {
		fmt.Printf("%d - %v\n", e.GetID(), *e.Name)
	}
	fmt.Println()
	return nil
}

func addUserToTeam(teamID int64, username string) error {

	// validation
	if teamID == 0 || username == "" {
		log.Fatal("required argument missing")
	}
	// end validation

	// username
	usr, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return err
	}
	// end user

	// team
	team, _, err := client.Teams.GetTeam(ctx, teamID)
	if err != nil {
		return err
	}
	// end team

	// prompt
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Are you sure you want to add '%s' (%s) to '%s' team?: [Y/n]",
		username, usr.GetName(), team.GetName())
	resp, _ := reader.ReadString('\n')
	if resp != "Y\n" {
		return nil
	}
	//end prompt

	// is already member
	isMember, _, err := client.Teams.IsTeamMember(ctx, teamID, username)
	if err != nil {
		return err
	}
	if isMember {
		fmt.Printf("%s already member of this team", username)
		return nil
	}
	// end if member

	// add user
	opt := &github.TeamAddTeamMembershipOptions{}
	_, _, err = client.Teams.AddTeamMembership(ctx, teamID, username, opt)
	if err != nil {
		return err
	}
	fmt.Printf("%s has been added to this team", username)
	// end add user

	return nil

}
