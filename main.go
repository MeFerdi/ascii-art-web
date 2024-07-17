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

	http.HandleFunc("/", web.Handler)
	http.HandleFunc("/ascii", web.AsciiArtHandler)
	http.HandleFunc("/ascii-art-live", web.AsciiArtLiveHandler)
	fmt.Println("Initializing server at http://localhost:8080")

	port := ":8080"
	http.ListenAndServe(port, nil)
}
