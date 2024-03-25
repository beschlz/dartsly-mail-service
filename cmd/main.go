package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello`))
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {

	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
