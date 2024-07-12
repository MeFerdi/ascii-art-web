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
	fmt.Println("Starting server at http://localhost:8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii", asciiArtHandler)
	http.HandleFunc("/ascii-art-live", asciiArtLiveHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "404.html")
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "500.html")
		return
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "405.html")
		return
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

	if str == "" || bannerStyle == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "400.html")
		return
	}

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "500.html")
		return
	}

	data := struct {
		Art string
	}{
		Art: art,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "500.html")
		return
	}
}

func asciiArtLiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "405.html")
		return
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

	if str == "" || bannerStyle == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.ServeFile(w, r, "400.html")
		return
	}

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "500.html")
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(art))
}
