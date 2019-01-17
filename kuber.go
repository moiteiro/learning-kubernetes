package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello there")
	})

	http.ListenAndServe("0.0.0.0:8080", nil)
}
