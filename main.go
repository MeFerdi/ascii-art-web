package main

import (
	"fmt"
	"log"
	"net/http"
	
)

func main() {
	// Create a file server for the "templates" directory
	fs := http.FileServer(http.Dir("templates"))

	// Handle requests starting with "/templates/" by stripping the prefix 
	// and serving files from "templates" directory
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	// Define a handler function for the root route ("/")
	handlerHome := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	}

	// Register the handler function for the root route ("/")
	http.HandleFunc("/", handlerHome)

	
	// Specify the port to listen on
	port := ":8080"

	// Start the HTTP server
	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
