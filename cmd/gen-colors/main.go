package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		infile string
		// width, height int
	)

	// flag.IntVar(&height, "height", 1000, "output height in pixels")
	flag.StringVar(&infile, "infile", "", "path to input file")

	flag.Parse()

	f, err := os.Open(infile)
	if err != nil {
		log.Println("Couldn't open infile")
		log.Fatal(err)
	}
	palettes, err := BuildPaletteFromFile(f)
	log.Printf("color palette: %#v", palettes)

}

type Palette struct {
	Name   string
	Colors []color.Color
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

		log.Printf("Record found: %#v\n", record)
		name, colors, err := ExtractRecord(record)
		if err != nil {
			return nil, fmt.Errorf("Error extracting palette from record: %w", err)
		}
		p = append(p, Palette{Name: name, Colors: colors})

	}

	return p, nil
}

func ExtractRecord(record []string) (string, []color.Color, error) {
	name := record[0]
	colors := make([]color.Color, 0)

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
