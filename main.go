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
	"os"
)

var usage = `Usage: pmt [options...] [FOLDER]
options:
    -c    Export converted palette as blank css
`

var (
	c = flag.Bool("c", true, "")
)

type Spread struct {
	Folder string   `json:"folder"`
	Images []string `json:"images"`
}

/* Layout ready image type
type photo struct {
	width  int
	height int
	format string
	pal    palette.Palette
} */

var subimages []image.Image // RGBA, etc. images

type sheet struct {
	name string
	w    int
	h    int
	pal  color.Palette
}

func main() {
}

func palFromImg(name string) (color.Palette, error) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("Couldn't open file")
	}
	im, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Couldn't decode test file")
	}

	q := MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, 256), im)
	if err != nil {
		fmt.Println("Couldn't quantize image")
	}

	return p, nil
}

func fetchPalettes(sheet []string) ([]MedianCutQuantizer, error) {
	sheet := []MedianCutQuantizer{}
	for _, n := range sheet {
		sheet = append(sheet, palFromImg(n))
	}
	if err != nil {
		fmt.Println("Couldn't create pallete from image")
		return
	}

	return sheet, nil
}

// Get images from dir
func fetchImages(folder string) ([]string, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	out := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		out = append(out, f.Name())
	}
	return out, nil
}

/* test if file has image extension
func isImage(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return (ext == ".gif" || ext == ".jpg" || ".jpeg" || ext == ".png")
} */
