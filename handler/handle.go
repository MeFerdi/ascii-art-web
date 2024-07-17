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
		w.WriteHeader(http.StatusNotFound)
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	str := r.FormValue("textData")
	bannerStyle := r.FormValue("banner")

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if len(str) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := struct {
		Art string
	}{
		Art: art,
	}

	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func AsciiArtLiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(art))
}
