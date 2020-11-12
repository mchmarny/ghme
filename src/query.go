package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	queryRepoListCommand = cli.Command{
		Name:  "query",
		Usage: "Queries local repository file for match",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "file",
				Usage: "Input file path",
			},
			cli.StringFlag{
				Name:  "value",
				Usage: "Value to query",
			},
		},
		Action: func(c *cli.Context) error {
			path := c.String("file")
			query := c.String("value")
			if path == "" || query == "" {
				return errors.New("either output path or query required")
			}
			return queryRepoList(path, query)
		},
	}
)

func queryRepoList(path, query string) error {
	items, err := queryList(path, query)
	if err != nil {
		return errors.Wrapf(err, "processing query: %s", query)
	}

	out, err := json.Marshal(&list{Items: items})
	if err != nil {
		return errors.Wrap(err, "error serializing content")
	}

	fmt.Println(string(out))
	return nil
}

func queryList(path, query string) (items []*item, err error) {
	if path == "" {
		return nil, errors.New("path required")
	}
	in, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading file: %s", path)
	}

	var l list
	if err := json.Unmarshal(in, &l); err != nil {
		return nil, errors.Wrapf(err, "error deserializing file: %s", path)
	}

	q := strings.TrimSpace(strings.ToLower(query))
	items = make([]*item, 0)
	for _, r := range l.Items {
		if contains(q, r.Arg, r.Subtitle, r.Title) {
			items = append(items, r)
		}
	}
	return
}

func contains(q string, s ...string) bool {
	for _, v := range s {
		if strings.Contains(strings.TrimSpace(strings.ToLower(v)), q) {
			return true
		}
	}
	return false
}
