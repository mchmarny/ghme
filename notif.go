package main

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/urfave/cli"
)

var (
	getNotificationsCommand = cli.Command{
		Name:  "notif",
		Usage: "Lists your notifications",
		Action: func(c *cli.Context) error {
			return getMyNotifications()
		},
	}
)

// GetMyNotifications prints your own notifications
func getMyNotifications() error {

	fmt.Println()

	opt := &github.NotificationListOptions{
		All:           true,
		Participating: true,
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	var allItems []*github.Notification
	for {
		list, resp, err := client.Activity.ListNotifications(ctx, opt)
		if err != nil {
			return err
		}
		allItems = append(allItems, list...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, n := range allItems {
		fmt.Printf("[%s] %s:%s - %s (%s)\n",
			*n.Repository.Name, *n.Subject.Type, *n.Reason, *n.Subject.Title, *n.URL)
	}
	fmt.Println()
	return nil

}
