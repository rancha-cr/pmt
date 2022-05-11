// PMT := "photo-mechanical transfer." Technology of paper-based typesetting using light-sensitive paper to carry images.
package main

import (
	"flag"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/ericpauley/go-quantize/quantize"
)

var usage = `Usage: pmt [options...] [FOLDER]
options:
    -c    Export converted palette as blank css
`

var (
	n = flag.Bool("n", asCSS, "")
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

func main() {
	r := 100
	return r
}

// Get images from dir
func loadDir(folder string) ([]string, error) {
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
