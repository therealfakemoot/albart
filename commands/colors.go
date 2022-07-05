package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"

	"github.com/therealfakemoot/albart"
)

var colorFlags = []cli.Flag{}

var ColorsCommand = &cli.Command{
	Name:  "colors",
	Usage: "list known colors",
	Flags: colorFlags,
	Action: func(ctx *cli.Context) error {
		app, err := albart.NewApp(ctx.String("profile"))
		if err != nil {
			return fmt.Errorf("error loading app: %w", err)
		}
		for pName := range app.Colors {
			fmt.Println(pName)
		}
		return nil
	},
}
