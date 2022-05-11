// PMT := "photo-mechanical transfer." Technology of paper-based typesetting using light-sensitive paper to carry images.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/mccutchen/palettor"
	"github.com/nfnt/resize"
)

type Sheet struct {
	Color  color.Color `json:"color"`
	Weight float64     `json:"weight"`
	Hex    string      `json:"hex"`
	PMS    string      `json:"pms"`
}

type Spread struct{}

func main() {
	var (
		k          = flag.Int("k", 6, "Palette size (default: 3")
		jsonOutput = flag.Bool("json", false, "Output as JSON")
	)
	var usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [INPUT]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	testImage := "img/illustrations/5.jpg"
	testPalette := palFromImage(testImage, 6)

	for _, col := range testPalette.Colors() {
		fmt.Printf("color: %v; weight: %v\n", col, testPalette.Weight(col))
	}
}

// for colors in palette, convert to hex (and pantone)
/* func palToHex(p palettor.Palette) (color.Color, float64) {
	// out := []string{}
	for _, color := range p.Colors() {
		return color, p.Weight(color)
	}
} */

func palFromImage(name string, k int) *palettor.Palette {
	file, err := os.Open(name)
	og, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Couldn't open file")
	}
	defer file.Close()

	// reduce image size
	im := resize.Thumbnail(200, 200, og, resize.Lanczos3)

	// k = num of most dominant colors
	maxIterations := 100 // num of iterations until halt
	pal, err := palettor.Extract(k, maxIterations, im)
	if err != nil {
		log.Fatalf("Image too small")
	}
	return pal
}

// Get images from dir
func fetchImages(folder string) []string {
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Couldn't open folder")
	}
	out := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		out = append(out, f.Name())
	}
	return out
}

// json file w/ palette and converted hex/pms codes
var page = `
`
