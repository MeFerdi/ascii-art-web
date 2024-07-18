package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handler"
)

func main() {
	// Serve Static files (HTML/CSS)
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			handler.Handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("404 error: %s", r.URL.Path)
		}
	})

	http.HandleFunc("/ascii", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ascii" {
			handler.AsciiArtHandler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("404 error: %s", r.URL.Path)
		}
	})

	port := ":8080"
	fmt.Printf("Starting Server at port %v  \n", port)
	http.ListenAndServe(port, nil)
}
