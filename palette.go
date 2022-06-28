package albart

import (
	"encoding/csv"
	"fmt"
	"image/color"
	"io"
	"strconv"
)

type Palette struct {
	Name   string
	Colors []color.RGBA
}

func BuildPaletteFromFile(r io.Reader) ([]Palette, error) {
	p := make([]Palette, 0)

	csvReader := csv.NewReader(r)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error reading record: %w", err)
		}

		name, colors, err := ExtractRecord(record)
		if err != nil {
			return nil, fmt.Errorf("Error extracting palette from record: %w", err)
		}
		p = append(p, Palette{Name: name, Colors: colors})

	}

	return p, nil
}

func ExtractRecord(record []string) (string, []color.RGBA, error) {
	name := record[0]
	colors := make([]color.RGBA, 0)

	for _, color := range record[1:] {
		c, err := extractRGB(color)
		if err != nil {
			return "", nil, err
		}
		colors = append(colors, c)
	}

	return name, colors, nil
}

func extractRGB(hex string) (color.RGBA, error) {
	rgb, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return color.RGBA{}, err
	}

	return color.RGBA{
		R: uint8((rgb >> 16)),
		G: uint8((rgb >> 8) & 0xFF),
		B: uint8((rgb & 0xFF)),
	}, nil
}
