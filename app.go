package albart

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
)

type App struct {
	Colors map[string][]color.RGBA
}

func (a *App) LoadPalettes(d string) error {

	palettes, err := filepath.Glob(filepath.Join(d, "*.csv"))
	if err != nil {
		return fmt.Errorf("error scanning for palettes: %w\n", err)
	}

	colorPalettes := make(map[string][]color.RGBA)
	swatchNames := make([]string, 0)
	for _, p := range palettes {
		f, err := os.Open(p)
		if err != nil {
			return fmt.Errorf("error opening palette %s: %w\n", p, err)
		}

		pal, err := BuildPaletteFromFile(f)
		if err != nil {
			return fmt.Errorf("error building palette %s: %w\n", p, err)
		}
		for _, swatch := range pal {
			colorPalettes[swatch.Name] = swatch.Colors
			swatchNames = append(swatchNames, swatch.Name)
		}
	}
	a.Colors = colorPalettes
	return nil
}
