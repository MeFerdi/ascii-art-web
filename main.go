package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	// Create a file server for the "templates" directory
	fs := http.FileServer(http.Dir("templates"))

	// Handle requests starting with "/templates/" by stripping the prefix
	// and serving files from "templates" directory
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	// Register the handler function for the root route ("/")
	http.HandleFunc("/", handlers.HandlerHome)

	port := ":8080"

	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
