package commands

import (
	"github.com/la3mmchen/heuer/internal/types"
	"github.com/urfave/cli"
)

// GetApp <tbd>
func GetApp(cfg types.Configuration, version string) *cli.App {
	app := cli.NewApp()
	app.Name = "heuer"
	app.Usage = "A simple cli to print cards of a defined trello list"
	app.Version = version

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print app version",
	}

	app.Commands = []cli.Command{
		read(cfg),
	}

	return app
}
