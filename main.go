package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a handler function for handling HTTP requests
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!") // Response to client
	}

	// Register the handler function for the root route ("/")
	http.HandleFunc("/", handler)

	// Specify the port to listen on
	port := ":8080" // Note the colon ":" before the port number

	// Start the HTTP server
	fmt.Printf("Server is listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
