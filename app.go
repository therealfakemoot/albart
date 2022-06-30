package albart

import (
	"fmt"
	"image/color"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type AppConfig struct {
	Styles   []string `toml:"general.styles"`
	Colors   []string `toml:"general.colors"`
	Width    int      `toml:"general.width"`
	Height   int      `toml:"general.height"`
	Outfile  string   `toml:"general.outfile"`
	ColorDir string   `toml:"general.colorDir"`
	Seed     int      `toml:"general.seed"`
}

type App struct {
	Colors map[string][]color.RGBA
	Styles []string // placeholder until i implement proper art style switching
}

func (a *App) LoadConfig(r io.Reader) error {
	var ac AppConfig

	_, err := toml.DecodeReader(r, &ac)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", ac)
	return nil
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
