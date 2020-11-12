package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

var (
	ctx = context.Background()

	// AppName is the name of the app
	AppName = "ghme"

	// AppVersion set in compile time
	AppVersion = "v0.0.1-default"
)

func main() {
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "prints only the version",
	}

	app := cli.NewApp()
	app.Name = AppName
	app.Usage = "GitHub helper utility"
	app.Compiled = time.Now()
	app.Version = AppVersion
	app.HideHelp = false
	app.Commands = []cli.Command{
		listRepoListCommand,
		queryRepoListCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\nError: %v\n\n", err)
	} else {
		fmt.Println()
	}
}
