package main

import (
	"net/http"
	"web/handler"
)

func main() {

	http.HandleFunc("/",handler.Handler)


	port := ":8080"
	http.ListenAndServe(port,nil)
}
