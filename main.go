package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	web "web/ascii" // Ensure this matches your module path
)

var templates = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func main() {
	// Create a file server for the "templates" directory
	fs := http.FileServer(http.Dir("templates"))

	// Handle requests starting with "/templates/" by stripping the prefix
	// and serving files from "templates" directory
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii", asciiArtHandler)
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

	if str == "" || bannerStyle == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Art string
	}{
		Art: art,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
