package albart

import (
	"fmt"
	"image/color"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type TomlEnvelope struct {
	AppConfig AppConfig `toml:"app"`
}

type AppColors struct {
	Dir       string   `toml:"dir"`
	Preferred []string `toml:"preferred"`
}

type StyleConfig map[string]map[string]interface{}

type AppConfig struct {
	Styles  StyleConfig `toml:"style"`
	Colors  AppColors   `toml:"colors"`
	Width   int         `toml:"width"`
	Height  int         `toml:"height"`
	Outfile string      `toml:"outfile"`
	Seed    int         `toml:"seed"`
}

func NewApp(s string) (*App, error) {
	var app App

	profileFile, err := os.Open(s)
	if err != nil {
		return nil, fmt.Errorf("error opening profile: %w", err)
	}
	err = app.LoadConfig(profileFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing profile: %w", err)
	}

	err = app.LoadPalettes(app.Conf.Colors.Dir)
	if err != nil {
		return nil, fmt.Errorf("error scanning for palettes: %w", err)
	}
	return &app, nil
}

type App struct {
	Conf   AppConfig
	Colors map[string][]color.RGBA
	Styles []string // placeholder until i implement proper art style switching
}

func (a *App) LoadConfig(r io.Reader) error {
	var te TomlEnvelope

	_, err := toml.DecodeReader(r, &te)
	if err != nil {
		return err
	}
	a.Conf = te.AppConfig
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
