package albart

import (
	// "fmt"
	"github.com/urfave/cli/v2"
)

var RootCommand *cli.App = &cli.App{
	Name:  "albart",
	Usage: "Procedurally generate cover art",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "profile",
			Aliases:  []string{"p"},
			Required: true,
			Usage:    "TOML file containing generation `PROFILE`",
		},
	},
	Action: func(*cli.Context) error {
		return nil
	},
}
