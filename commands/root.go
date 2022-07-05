package commands

import (
	"github.com/urfave/cli/v2"
)

var RootCommand *cli.App = &cli.App{
	Name:  "albart",
	Usage: "Procedurally generate cover art",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "profile",
			Usage:    "TOML file containing render parameters",
			Required: true,
		},
	},
	Commands: []*cli.Command{
		ColorsCommand,
		RenderCommand,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
