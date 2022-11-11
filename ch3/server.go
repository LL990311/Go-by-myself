package main

import (
	"Go-by-myself/ch3/mandelbrot"
	"Go-by-myself/ch3/surface"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	//e3.4 upload surface on the localhost:8080 server
	//e3.9 mandelbrot server
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := request.FormValue("type")
		if t != "" {
			_t, err := strconv.Atoi(t)
			if err != nil {
				fmt.Fprintf(writer, "Invalid input size %s", err)
			}

			if _t == 1 {
				mandelbrot.Mandelbrot(writer)
			} else {
				mandelbrot.E3(writer)
			}

		}
		writer.Header().Set("Content-Type", "image/svg+xml")
		surface.Surface(writer)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
