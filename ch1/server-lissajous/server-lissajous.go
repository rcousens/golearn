package main

import (
	"net/http"
	"log"
	"image/gif"
	"image"
	"math/rand"
	"math"
	"image/color"
	"strconv"
	"fmt"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	options := make(map[string]interface{})

	options["cycles"] = 5
	options["res"] = 0.001
	options["size"] = 100
	options["nframes"] = 64
	options["delay"] = 8


	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		if k == "res" {
			options[k], _ = strconv.ParseFloat(v[0], 64)
		} else {
			options[k], _ = strconv.Atoi(v[0])
		}

	}

	freq := rand.Float64() * 3.0

	anim := gif.GIF{LoopCount: options["nframes"].(int)}

	phase := 0.0

	for i := 0; i < options["nframes"].(int); i++ {
		rect := image.Rect(0, 0, 2*options["size"].(int)+1, 2*options["size"].(int)+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(options["cycles"].(int))*2*math.Pi; t += options["res"].(float64) {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(options["size"].(int) + int(x*float64(options["size"].(int))+0.5), options["size"].(int)+int(y*float64(options["size"].(int))+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, options["delay"].(int))
		anim.Image = append(anim.Image, img)
	}

	fmt.Printf("%v\n", options)
	gif.EncodeAll(w, &anim)
}