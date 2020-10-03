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

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "print app version",
	}

	app.Commands = []cli.Command{
		read(cfg),
	}

	return app
}
