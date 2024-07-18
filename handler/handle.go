package web

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	web "web/ascii"
)

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseFiles(filepath.Join("templates", "index.html"))
	if err != nil {
		log.Printf("404 Not Found: %v", err)
		templates = nil // Set templates to nil to indicate the template couldn't be loaded
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if templates == nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template %s: %v", tmpl, err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if templates == nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template index.html: %v", err)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	str := r.FormValue("textData")
	bannerStyle := r.FormValue("banner")

	if len(str) == 0 {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	art, err := web.PrintAscii(str, bannerStyle)
	if err != nil {
		
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error generating ASCII art: %v", err)
		return
	}

	data := struct {
		Art string
	}{
		Art: art,
	}

	renderTemplate(w, "index", data)
}