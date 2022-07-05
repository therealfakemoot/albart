package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"

	"github.com/therealfakemoot/albart"
)

var renderFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     "track",
		Usage:    "Name of media item to generate cover art for",
		Required: true,
	},
}

var RenderCommand = &cli.Command{
	Name:  "render",
	Usage: "render cover art",
	Flags: renderFlags,
	Action: func(ctx *cli.Context) error {
		var app albart.App
		err := app.LoadPalettes(ctx.String("colorDir"))
		if err != nil {
			return fmt.Errorf("error scanning for palettes: %w", err)
		}

		return nil
	},
}
