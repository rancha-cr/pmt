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
		return
	}
	i, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Couldn't decode file")
		return
	}
	q := MedianCutQuantizer{}
	p := q.Quantize(make([]color.Color, 0, 256), i)
	fmt.Println(p)
	return p
}

// Quantize using median cut method
