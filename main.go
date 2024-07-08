package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Serve static files (HTML, CSS)
    fs := http.FileServer(http.Dir("templates"))
    http.Handle("/templates/", http.StripPrefix("/templates/", fs))

    // Define a handler function for handling HTTP requests
    handler := func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "templates/index.html")
    }

    // Register the handler function for the root route ("/")
    http.HandleFunc("/", handler)

    // Specify the port to listen on
    port := ":8080" // Note the colon ":" before the port number

    // Start the HTTP server
    fmt.Printf("Server is listening on port %s...\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}
