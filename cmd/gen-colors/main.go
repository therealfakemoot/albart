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

func buildPalette(infile string) []color.Color {
	p := make([]color.Color, 0)
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
			log.Fatal(err)
		}

		log.Printf("Record found: %#v\n", record)
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
	rgb, err := strconv.ParseUint(hex, 16, 8)
	if err != nil {
		return color.RGBA{}, err
	}

	blueMask := uint64(0xFF0000)
	greenMask := uint64(0xFF00)
	redMask := uint64(0xFF)

	return color.RGBA{
		R: uint8((rgb & redMask)),
		G: uint8((rgb & blueMask)),
		B: uint8((rgb & greenMask)),
	}, nil
}
