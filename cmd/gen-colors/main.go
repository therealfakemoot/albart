package main

import (
	"encoding/csv"
	"flag"
	"image/color"
	"io"
	"log"
	"os"
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

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// log.Println(record)
	}

}

func buildPalette() *color.Color {

}
