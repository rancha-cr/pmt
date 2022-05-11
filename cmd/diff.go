package diff

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

// Quuantize using median cut method
func Quant(f string) color.Palette {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Couldn't open file")
	}

	i, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Couldn't decode file")
	}

	q := MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	return p
}

func Diff() {
}

// Convert RGB color to Pantone matching system (PMS)
