package main

import (
	"flag"

	"github.com/fogleman/gg"

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
	dc := gg.NewContext(width, height)

	dc.SaveFaile(outfile)
}
