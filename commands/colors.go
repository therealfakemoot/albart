package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"

	"github.com/therealfakemoot/albart"
)

var colorFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "colorDir",
		Usage:    "Directory containing color swatch CSV files",
		Required: true,
	},
}

var ColorsCommand = &cli.Command{
	Name:  "colors",
	Usage: "list known colors",
	Flags: colorFlags,
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
}
