package main

import (
	"fmt"
	"net/http"

	web "web/handler"
)

func main() {
	// Serve Static files (HTML/CSS)
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			web.Handler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/ascii", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ascii" {
			web.AsciiArtHandler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/ascii-art-live", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/ascii-art-live" {
			web.AsciiArtLiveHandler(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	port := ":8080"
	fmt.Printf("Starting Server at port %v  \n", port)
	http.ListenAndServe(port, nil)
}
