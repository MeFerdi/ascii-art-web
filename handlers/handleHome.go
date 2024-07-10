package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// Serve the index.html file

var templates = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "templates/404.html")
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.ServeFile(w, r, "templates/500.html")
		return
	}
}
