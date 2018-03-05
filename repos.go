package main

import (
	"github.com/google/go-github/github"
)

// PrintRepos prints repos for a given org
func PrintRepos(client *github.Client, org string) {
	logger.Printf("Listing repos for %s", org)
	repos, err := ListRepos(client, org)
	if err != nil {
		logger.Fatalf("Error on repo list: %v", err)
	}
	for i, e := range repos {
		logger.Printf("%d: %v", i, *e.FullName)
	}
}

// ListRepos lists repos for specific org
func ListRepos(client *github.Client, org string) (repos []*github.Repository, err error) {

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var allItems []*github.Repository
	for {
		list, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
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
