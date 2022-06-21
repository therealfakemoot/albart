package main

import (
	"encoding/csv"
	"flag"
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

	log.Printf("color palette: %#v", buildPalette(infile))

}

type Palette struct {
	Name   string
	Colors []color.Color
}

func buildPalette(infile string) []Palette {
	p := make([]Palette, 0)
	f, err := os.Open(infile)
	if err != nil {
		log.Println("Couldn't open infile")
		log.Fatal(err)
	}

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading record: %s\n", err)
		}

		log.Printf("Record found: %#v\n", record)
		name, colors, err := ExtractRecord(record)
		if err != nil {
			log.Fatalf("Error extracting palette from record: %s\n", err)
		}
		p = append(p, Palette{Name: name, Colors: colors})

	}

	return p
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
