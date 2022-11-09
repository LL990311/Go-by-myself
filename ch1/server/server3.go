package main

import (
	"fmt"
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

var palette = []color.Color{color.Black, color.RGBA{R: 10, G: 249, B: 84, A: 1}, color.White}

const (
	cycle   = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {

	http.HandleFunc("/handler", handler3)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycleStr := r.FormValue("cycles")
		if cycleStr == "" {
			lissajous(w, cycle, 1)
		}
		userCycle, err := strconv.Atoi(cycleStr)
		if err != nil {
			fmt.Fprintf(w, "Invalid input cycle %s", err)
		}
		lissajous(w, float64(userCycle), 2)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func lissajous(out io.Writer, userCycle float64, index uint8) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < userCycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
