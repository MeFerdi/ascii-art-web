package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"

    "ascii-art-web/art"
)

type PageData struct {
    ASCIIArt string
}

func main() {
    // Create a file server for the "templates" directory
    fs := http.FileServer(http.Dir("templates"))

    // Handle requests starting with "/templates/" by stripping the prefix
    // and serving files from "templates" directory
    http.Handle("/templates/", http.StripPrefix("/templates/", fs))

    // Register the handler function for the root route ("/")
    http.HandleFunc("/", handlerHome)

    // Specify the port to listen on
    port := ":8080"

    // Start the HTTP server
    fmt.Printf("Server is listening on port %s...\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

func handlerHome(w http.ResponseWriter, r *http.Request) {
    // Serve the index.html file
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    // Check if the method is POST (form submission)
    if r.Method == http.MethodPost {
        // Parse form data
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Unable to parse form", http.StatusBadRequest)
            return
        }

        // Retrieve values from form
        bannerType := r.Form.Get("banner")   // Radio button value
        printText := r.Form.Get("printText") // Textarea value

        if bannerType == "standard" {
            bannerType = "standard.txt"
        } else if bannerType == "shadow" {
            bannerType = "shadow.txt"
        } else if bannerType == "thinkertoy" {
            bannerType = "thinkertoy.txt"
        }

        // Call your function to read the ASCII art file and generate the art
        fileSlice, err := art.Reading(bannerType)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error reading ASCII art file: %v", err), http.StatusInternalServerError)
            return
        }

        // Generate ASCII art based on the file slice and user input text
        asciiArt := art.PrintArt(fileSlice, printText)

        // Create a PageData struct instance with ASCII art
        data := PageData{
            ASCIIArt: asciiArt,
        }

        // Parse the HTML template
        tmpl, err := template.ParseFiles("templates/index.html")
        if err != nil {
            http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
            return
        }

        // Execute the template with data and write to response
        err = tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
            return
        }

        return
    }

    // For GET requests or initial load, serve the index.html file
    http.ServeFile(w, r, "templates/index.html")
}
