package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

type config struct {
	cycles,
	res float64
	size,
	nframes,
	delay int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		conf := config{5, 0.001, 100, 64, 8}
		for k, v := range r.Form {
			val, err := strconv.Atoi(v[0])
			if err != nil {
				continue
			}
			switch k {
			case "cycles":
				conf.cycles = float64(val)
			case "res":
				conf.cycles = float64(val)
			case "size":
				conf.size = val
			case "nframe":
				conf.nframes = val
			case "delay":
				conf.delay = val
			}
		}
		lissajous(w, conf)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, conf config) {
	var (
		cycles  = conf.cycles  // number of complete x oscillator revolutions
		res     = conf.res     // angular resolution
		size    = conf.size    // image canvas covers [-size..+size]
		nframes = conf.nframes // number of animation frames
		delay   = conf.delay   // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
