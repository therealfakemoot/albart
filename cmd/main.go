package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fogleman/gg"
	"github.com/nfnt/resize"
	"github.com/therealfakemoot/albart"
)

func main() {
	var (
		logo          string
		outfile       string
		track         string
		width, height int
	)

	flag.IntVar(&width, "width", 1000, "output width in pixels")
	flag.IntVar(&height, "height", 1000, "output height in pixels")
	flag.StringVar(&logo, "logo", "", "path to logo png")
	flag.StringVar(&outfile, "outfile", "", "path to output file")
	flag.StringVar(&track, "track", "", "track title")

	flag.Parse()

	if outfile == "" {
		fmt.Println("please provide the -outfile flag")
		log.Fatal()
	}

	logoFile, err := gg.LoadPNG(logo)
	if err != nil {
		log.Fatalf("error opening logo: %s\n", err)
	}
	// logoContext := gg.NewContext(
	logoContext := gg.NewContextForImage(resize.Resize(uint(width/4), 0, logoFile, resize.Lanczos3))

	baseContext := gg.NewContext(width, height)

	baseContext.SetRGB(0, 0, 0)
	baseContext.Clear()
	baseContext.SetRGB(1, 1, 1)
	// dc.MeasureString() for adjusting text rendering
	// DrawRoundedRectangle might be nice
	corners := albart.LaneCorners(albart.North, width, height, logoContext.Width(), logoContext.Height())

	anchor := albart.LaneAnchor(corners)

	baseContext.DrawImageAnchored(logoContext.Image(), int(anchor.X), int(anchor.Y), 0, 0)

	// need to rescale this to desired output dimensions
	err = baseContext.SavePNG(outfile)
	if err != nil {
		log.Fatalf("error saving outfile: %s\n", err)
	}
}
