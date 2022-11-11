package main

import (
	"Go-by-myself/ch3/surface"
	"log"
	"net/http"
)

func main() {

	//e3.4 upload surface on the localhost:8080 server
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		surface.Surface(writer)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
