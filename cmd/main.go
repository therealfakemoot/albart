package main

import (
	// "flag"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	// "path/filepath"
	"time"

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
	// I could use a closure here to inject state into the RootCommand
	// but actually, the root command isn't gonna do much, i'm gonna have
	// subcommands for all the actual work
	// albart.RootCommand.Action =
	if err := albart.RootCommand.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	return

	rand.Seed(time.Now().UnixNano())
	//rand.Seed(8675309)
	var (
		logo    string
		outfile string
		// track         string
		// style         string
		profile    string
		paletteDir string
		// palette       string
		listPalettes  bool
		width, height int
	)

	var app albart.App
	err := app.LoadPalettes(paletteDir)
	if err != nil {
		log.Fatalf("error scanning for palettes: %s\n", err)
	}

	profileFile, err := os.Open(profile)
	if err != nil {
		log.Fatalf("Error opening profile: %s\n", err)
	}
	err = app.LoadConfig(profileFile)
	if err != nil {
		log.Fatalf("Error parsing profile: %s\n", err)
	}

	return
	// below here is the old flow

	if listPalettes {
		for pName := range app.Colors {
			fmt.Println(pName)
		}
		// fmt.Printf("%#v\n", app.Colors)
		return
	}

	if outfile == "" {
		fmt.Println("please provide the -outfile flag")
		log.Fatal()
	}

	logoFile, err := gg.LoadPNG(logo)
	if err != nil {
		log.Fatalf("error opening logo: %s\n", err)
	}
	// logoContext := gg.NewContext(
	logoContext := gg.NewContextForImage(resize.Resize(uint(float64(width)*0.45), 0, logoFile, resize.Lanczos3))

	/*
		rand.Shuffle(len(swatchNames), func(i, j int) {
			swatchNames[i], swatchNames[j] = swatchNames[j], swatchNames[i]
		})
	*/

	c := generativeart.NewCanva(width, height)
	c.SetBackground(color.RGBA{0x00, 0x00, 0x00, 0xFF})
	c.FillBackground()

	/*
		c.SetColorSchema(colorPalettes[swatchNames[0]])
		if palette != "" {
			c.SetColorSchema(colorPalettes[palette])
		}
	*/
	c.Draw(arts.NewNoiseLine(1000))

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

	log.Println(app)
}
