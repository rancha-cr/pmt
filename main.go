// PMT := "photo-mechanical transfer." Technology of paper-based typesetting using light-sensitive paper to carry images.
package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	diff "andygo/cmd"

	"github.com/ericpauley/go-quantize/quantize"
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

// Get images from dir
func getImages(folder string) ([]string, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	out := []string{}
	for _, f := range files {
		if f.IsDir() || !isImage(f.Name()) {
			continue
		}
		out = append(out, f.Name())
	}
	quantize.MedianCutQuantizer(out)
	return out, nil
}

func isImage(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return (ext == ".gif" || ext == ".jpg" || ".jpeg" || ext == ".png")
}

func getPalette(file string) *image.Image {
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	p, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	return &p
}

func main() {
	f := "penis"
	diff.Quant(f)
}
