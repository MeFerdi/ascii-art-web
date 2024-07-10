package handlers

import (
	"net/http"
)

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	// Serve the index.html file
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// For GET requests or initial load, serve the index.html file
	http.ServeFile(w, r, "templates/index.html")
}
