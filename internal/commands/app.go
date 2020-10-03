package commands

import (
	"github.com/la3mmchen/treta/internal/types"
	"github.com/urfave/cli"
)

var (
	// TrelloToken <tbd>
	TrelloToken string
	// TrelloAppKey <tbd>
	TrelloAppKey string
	// List <tbd>
	List string
)

// GetApp <tbd>
func GetApp(cfg types.Configuration) *cli.App {
	app := cli.NewApp()
	app.Name = cfg.AppName
	app.Usage = cfg.AppUsage
	app.Version = cfg.AppVersion
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token, t",
			Destination: &TrelloToken,
			Value:       "",
			Usage:       "TrelloToken to authenticate.",
		},
		cli.StringFlag{
			Name:        "user, u",
			Destination: &TrelloAppKey,
			Value:       "",
			Usage:       "TrelloAppKey",
		},
		cli.StringFlag{
			Name:        "list, l",
			Destination: &List,
			Value:       cfg.TrelloList,
			Usage:       "List",
		},
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print app version",
	}

	app.Commands = []cli.Command{
		list(cfg),
	}

	return app
}
