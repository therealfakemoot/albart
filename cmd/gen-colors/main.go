package main

import (
	"flag"
	"log"
	"os"

	"github.com/therealfakemoot/albart"
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
	palettes, err := albart.BuildPaletteFromFile(f)
	log.Printf("color palette: %#v", palettes)

}
