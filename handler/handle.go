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

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "405.html")
		return
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

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

func AsciiArtLiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.ServeFile(w, r, "405.html")
		return
	}

	str := r.FormValue("text")
	bannerStyle := r.FormValue("banner")

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "500.html")
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(art))
}
