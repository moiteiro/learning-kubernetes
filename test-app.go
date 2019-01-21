package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Processing request")
		io.WriteString(w, "Hello there, updating content without destroying service")
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
