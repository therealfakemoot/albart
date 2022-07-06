package commands

import (
	"fmt"

	"github.com/therealfakemoot/albart"
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
	Before: func(ctx *cli.Context) error {

		app, err := albart.NewApp(ctx.String("profile"))
		if err != nil {
			return fmt.Errorf("error loading app: %w", err)
		}
		ctx.App.Metadata["app"] = app
		return nil
	},
}
