package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"

	"github.com/therealfakemoot/albart"
)

var colorFlags = []cli.Flag{}

var ColorsCommand = &cli.Command{
	Name:  "colors",
	Usage: "list known colors",
	Flags: colorFlags,
	Action: func(ctx *cli.Context) error {
		var app albart.App
		log.Printf("profile arg: %q", ctx.String("profile"))

		profileFile, err := os.Open(ctx.String("profile"))
		if err != nil {
			return fmt.Errorf("error opening profile: %w", err)
		}
		err = app.LoadConfig(profileFile)
		if err != nil {
			return fmt.Errorf("error parsing profile: %w", err)
		}

		err = app.LoadPalettes(app.Conf.Colors.Dir)
		if err != nil {
			return fmt.Errorf("error scanning for palettes: %w", err)
		}
		for pName := range app.Colors {
			fmt.Println(pName)
		}
		return nil
	},
}
