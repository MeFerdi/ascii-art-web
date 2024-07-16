package web

import (
	"net/http"
	"path/filepath"
	"text/template"

	web "web/ascii"
)

var templates = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	str := r.FormValue("textData")
	bannerStyle := r.FormValue("banner")

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Art string
	}{
		Art: art,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AsciiArtLiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(art))
}
