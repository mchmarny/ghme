package main

import (
	"github.com/google/go-github/github"
)

// PrintTeams prints teams and its members
func PrintTeams(client *github.Client, org string) {
	logger.Printf("Listing teams and it's members for %s", org)
	teams, err := ListTeams(client, org)
	if err != nil {
		logger.Fatalf("Error on team list: %v", err)
	}
	for i, e := range teams {
		logger.Printf("%d: %v", i, *e.Name)
		users, err := ListTeamMembers(client, *e.ID)
		if err != nil {
			logger.Fatalf("Error on team list: %v", err)
		}
		for j, u := range users {
			logger.Printf("%d: %v", j, *u.Login)
		}
	}
}

// ListTeams lists teams for specific org
func ListTeams(client *github.Client, org string) (teams []*github.Team, err error) {

	opt := &github.ListOptions{
		PerPage: 10,
	}

	var allItems []*github.Team
	for {
		list, resp, err := client.Organizations.ListTeams(ctx, org, opt)
		if err != nil {
			return allItems, err
		}
		allItems = append(allItems, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allItems, nil

}

// ListTeamMembers lists all users for given team
func ListTeamMembers(client *github.Client, teamID int64) (teams []*github.User, err error) {

	opt := &github.OrganizationListTeamMembersOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allItems []*github.User
	for {
		list, resp, err := client.Organizations.ListTeamMembers(ctx, teamID, opt)
		if err != nil {
			logger.Fatal(err)
			return allItems, err
		}
		allItems = append(allItems, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return allItems, nil

}
