package commands

import (
	"github.com/urfave/cli/v2"
)

var RootCommand *cli.App = &cli.App{
	Name:  "albart",
	Usage: "Procedurally generate cover art",
	Commands: []*cli.Command{
		ColorsCommand,
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
