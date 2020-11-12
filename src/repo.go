package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
)

type list struct {
	Items []*item `json:"items"`
}

type item struct {
	Arg      string `json:"arg"`
	UID      int64  `json:"uid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

var (
	listRepoListCommand = cli.Command{
		Name:  "load",
		Usage: "Loads repositories accessable for GitHub user into local file",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "token",
				Usage: "GitHub Access Token",
			},
			cli.StringFlag{
				Name:  "file",
				Usage: "Output path path",
			},
		},
		Action: func(c *cli.Context) error {
			client, err := getGitHubClient(c.String("token"))
			if err != nil {
				return errors.Wrap(err, "error initializing GitHub client")
			}
			path := c.String("file")
			if path == "" {
				return errors.New("output path required")
			}
			return writeRepoList(client, path)
		},
	}
)

func getGitHubClient(token string) (client *github.Client, err error) {
	if token == "" {
		return nil, errors.New("access token required")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client = github.NewClient(oauth2.NewClient(ctx, ts))
	return
}

func writeRepoList(client *github.Client, path string) error {
	if client == nil {
		return errors.New("requires GitHub client")
	}
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error creating file")
	}
	defer f.Close()

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}
	var allItems []*github.Repository
	for {

		list, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			return err
		}
		allItems = append(allItems, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	items := make([]*item, 0)
	for _, r := range allItems {
		i := &item{
			Arg:      r.GetHTMLURL(),
			UID:      r.GetID(),
			Title:    r.GetName(),
			Subtitle: r.GetHTMLURL(),
		}
		items = append(items, i)
	}
	list := list{Items: items}

	b, err := json.Marshal(list)
	if err != nil {
		return errors.Wrap(err, "error mashalling list")
	}

	if _, e := f.Write(b); e != nil {
		return errors.Wrap(err, "error writting list")
	}

	fmt.Printf("saved %d repos", len(items))
	return nil
}
