package main

import (
	"net/http"

	"web/handler"
)

func main() {
	// Serve Static files (HTML/CSS)
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	http.HandleFunc("/", handler.Handler)

	port := ":8080"
	http.ListenAndServe(port, nil)
}
