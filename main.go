package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	web "web/ascii"
)

var templates = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func main() {
	// fmt.Println("Starting server on :8080...")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii", asciiArtHandler)
	http.HandleFunc("/ascii-art-live", asciiArtLiveHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
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

func asciiArtLiveHandler(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(art))
}
