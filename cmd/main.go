package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	// "math/rand"

	"github.com/fogleman/gg"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/nfnt/resize"

	"github.com/therealfakemoot/albart"
)

/*
func LogoAnchor(c *gg.Context, logo *gg.Context) gg.Point {
	rand.Seed(8675309)
	x, y := c.Width(), c.Height()
	w, h := logo.Width(), logo.Height()

	dir := []string{"N", "S", "E", "W"}
	switch dir[rand.Int31n(int32(len(dir)))] {
	case "W":
		// xSample := rand.Int63n(int64(0.5*w) + 1)
		bumper := .5 * w
		xSample := rand.Int63n(int64(x)) + bumper
	}

	// xSample := rand.Float64() * x
	// ySample := rand.Float64() * y
}
*/

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

	colors := []color.RGBA{
		{0x06, 0x7B, 0xC2, 0xFF},
		{0x84, 0xBC, 0xDA, 0xFF},
		{0xEC, 0xC3, 0x0B, 0xFF},
		{0xF3, 0x77, 0x48, 0xFF},
		{0xD5, 0x60, 0x62, 0xFF},
	}
	c := generativeart.NewCanva(width, height)
	c.SetBackground(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewNoiseLine(1000))
	// c.ToPNG("fuck.png")

	baseContext := gg.NewContextForRGBA(c.Img())
	// dc.MeasureString() for adjusting text rendering
	// DrawRoundedRectangle might be nice
	corners := albart.LaneCorners(albart.North, width, height, logoContext.Width(), logoContext.Height())

	anchor := albart.LaneAnchor(corners)

	// DrawImageAnchored draws the specified image at the specified anchor point. The anchor point is x - w * ax, y - h * ay, where w, h is the size of the image. Use ax=0.5, ay=0.5 to center the image at the specified point.
	baseContext.DrawImageAnchored(logoContext.Image(), int(anchor.X), int(anchor.Y), .5, .5)

	// need to rescale this to desired output dimensions
	err = baseContext.SavePNG(outfile)
	if err != nil {
		log.Fatalf("error saving outfile: %s\n", err)
	}
}
