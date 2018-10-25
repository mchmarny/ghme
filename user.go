package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

func printUser(org, username string) error {

	if username == "" {
		return fmt.Errorf("user argument required")
	}

	fmt.Printf("\nGetting user: %s\n\n", username)
	usr, err := getUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("ID: %d\n", usr.ID)
	fmt.Printf("Name: %s\n", usr.GetName())
	fmt.Printf("Login: %s\n", usr.GetLogin())
	fmt.Printf("Email: %s\n", usr.GetEmail())
	fmt.Printf("Location: %s\n", usr.GetLocation())
	fmt.Printf("Created: %v\n", usr.GetCreatedAt())
	fmt.Printf("Company: %s\n", usr.GetCompany())

	fmt.Println()

	return nil
}

func getUser(username string) (usr *github.User, err error) {

	if username == "" {
		return nil, fmt.Errorf("user argument required")
	}

	fmt.Printf("\nGetting user: %s\n", username)
	u, _, e := client.Users.Get(ctx, username)

	return u, e
}

func getUserByID(id int64) (usr *github.User, err error) {

	fmt.Printf("\nGetting user for ID: %d\n", id)
	u, _, e := client.Users.GetByID(ctx, id)

	return u, e
}

func getUserOrgActivity(org, usr string) (list []*github.Event, err error) {

	fmt.Printf("\nGetting user %s activity for %s\n", org, usr)

	opt := &github.ListOptions{PerPage: 10}
	var events []*github.Event
	for {
		list, resp, err := client.Activity.ListUserEventsForOrganization(
			ctx, org, usr, opt)
		if err != nil {
			return nil, err
		}
		events = append(events, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return events, nil
}
