// PMT := "photo-mechanical transfer." Technology of paper-based typesetting using light-sensitive paper to carry images.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mccutchen/palettor"
	"github.com/nfnt/resize"
)

type Swatch struct {
	Pantone string `json:"pantone"`
	PMS     string `json:"pms"`
	Hex     string `json:"hex"`
}

/* type Sheet struct {
	Filepath string      `json:"color"`
	Color    color.Color `json:"color"`
	Weight   float64     `json:"weight"`
	Hex      string      `json:"hex"`
	PMS      string      `json:"pms"`
} */

func main() {
	var (
		k = flag.Int("k", 6, "Palette size (default: 6")
		//		jsonOutput = flag.Bool("json", false, "Output as JSON")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [INPUT]\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	inputPath := "img/illustrations/4.jpg"
	if flag.NArg() > 0 {
		inputPath = flag.Args()[0]
	}

	inp, err := filepath.Abs(inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "filepath error")
		return
	}
	/*  src, err := os.Stat(folder)
	    if err != nil {
	        fmt.Fprintln(os.Stderr, "directory does not exist")
	        return
	    }
	    if !src.IsDir() {
	        fmt.Fprintln(os.Stderr, "path is not a directory")
	        return
	    } */
	num := *k
	if num < 0 {
		fmt.Fprintln(os.Stderr, "k cannot be smaller than 0")
		return
	}

	easel := palFromImage(inp, 6)
	colorList := []string{}
	for _, s := range easel.Colors() {
		colorList = append(colorList, convertColors(s))
	}

	spread := struct {
		Image  string `json:"image"`
		Color1 string `json:"color1"`
		Color2 string `json:"color2"`
		Color3 string `json:"color3"`
		Color4 string `json:"color4"`
		Color5 string `json:"color5"`
		Color6 string `json:"color6"`
	}{
		filepath.Base(inp),
		colorList[0],
		colorList[1],
		colorList[2],
		colorList[3],
		colorList[4],
		colorList[5],
	}

	http.HandleFunc("/i/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(inp, r.URL.Path[3:]))
	})
	http.HandleFunc("/data.json", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(spread)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
	})

	fmt.Print("Serving on http://localhost:2222")
	http.ListenAndServe(":2222", nil)
}

// for colors in palette, convert to hex (and pantone)
func convertColors(c color.Color) string {
	s := Swatch{}
	r, g, b, a := c.RGBA()

	s.Pantone = ""
	s.PMS = ""
	s.Hex = fmt.Sprintf("#%02x%02x%02x", r, g, b, a)

	return s.Hex
}

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

// html collage page
var page = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Collage</title>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
		<script type="text/javascript">
			$(function(){
				$.getJSON( "data.json", function(data) {
					document.title = "PMT • ";
					for (image of data.image) {
						$("#images").append("<img src='/i/"+image+"'></img>");
					}
					$images = $("#image img");
					$section = $("#images");
					$("#images img").click(function() {
						if ($section.hasClass("dimmed")) {
							$section.removeClass("dimmed");
							$images.removeClass("dim");
						} else {
							$section.addClass("dimmed");
							$images.not($(this)).addClass("dim");
						}
					});
				});
			});
		</script>
		<style type="text/css">
			#images {
				line-height: 0;
				-webkit-column-count: 3;
				-webkit-column-gap:   10px;
				-moz-column-count:    3;
				-moz-column-gap:      10px;
				column-count:         3;
				column-gap:           10px;  
			}
			#images img {
				max-width: 100%;
				height: auto;
				margin-bottom: 10px;
			}
			#images img.dim {
				opacity: 0.1;
			}
            #col1 {
                border: 15px solid data.color1
                font-size: 36px;
            }
            #col1 {
                border: 15px solid data.color2
                font-size: 36px;
            }
            #col1 {
                border: 15px solid data.color3
                font-size: 36px;
            }
            #col1 {
                border: 15px solid data.color4
                font-size: 36px;
            }
            #col1 {
                border: 15px solid data.color5
                font-size: 36px;
            }
            #col1 {
                border: 15px solid data.color6
                font-size: 36px;
            }
		</style>
	</head>
	<body>
		<section id="images">
		<div id="col1">1</div>
		<div id="col2">2</div>
		<div id="col3">3</div>
		<div id="col4">4</div>
		<div id="col5">5</div>
		<div id="col6">6</div>
		</section>
	</body>
</html>`
