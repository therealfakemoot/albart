package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"

	"github.com/therealfakemoot/albart"
)

var RootCommand *cli.App = &cli.App{
	Name:  "albart",
	Usage: "Procedurally generate cover art",
	Commands: []*cli.Command{
		&cli.Command{
			Name:  "colors",
			Usage: "list known colors",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "colorDir",
					Usage:    "Directory containing color swatch CSV files",
					Required: true,
				},
			},
			Action: func(ctx *cli.Context) error {
				var app albart.App
				err := app.LoadPalettes(ctx.String("colorDir"))
				if err != nil {
					return fmt.Errorf("error scanning for palettes: %w\n", err)
				}
				for pName := range app.Colors {
					fmt.Println(pName)
				}
				return nil
			},
		},
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
