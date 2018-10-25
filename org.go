package main

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

var (
	listOrgMemberCommand = cli.Command{
		Name:  "list",
		Usage: "List members of GitHub org",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "org, o",
				Usage: "GitHub Org ID",
			},
		},
		Action: func(c *cli.Context) error {

			org := c.String("org")

			if org == "" {
				return fmt.Errorf("org argument required")
			}

			return printOrgMembers(org)
		},
	}
)

func printOrgMembers(org string) error {

	if org == "" {
		return fmt.Errorf("org argument required")
	}

	fmt.Println()

	opt := &github.ListMembersOptions{
		PublicOnly:  false,
		Role:        "member",
		ListOptions: github.ListOptions{PerPage: 10},
	}
	var allItems []*github.User
	for {

		list, resp, err := client.Organizations.ListMembers(ctx, org, opt)
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
		usr, err := getUserByID(e.GetID())
		if err != nil {
			fmt.Printf("Error while getting user ID:%d - %v\n",
				e.GetID(), err)
		} else {
			fmt.Printf("%d - %s (%s <%s>)\n",
				usr.GetID(), usr.GetLogin(), usr.GetName(), usr.GetEmail())

			act, er := getUserOrgActivity(org, usr.GetLogin())
			if er != nil {
				fmt.Printf("Error while getting user activity for %s - %v\n",
					e.GetName(), er)
			} else {
				for i, a := range act {
					fmt.Printf("[%d] %s - %s - %s\n",
						i, a.CreatedAt, *a.Repo.Name, *a.Type)
				}
			}
		}
	}
	fmt.Println()
	return nil
}
