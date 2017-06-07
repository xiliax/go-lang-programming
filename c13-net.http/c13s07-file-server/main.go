package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	// Simple static webserver:
	//http.Handle("/static/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client"))))
	http.Handle("/", http.FileServer(http.Dir("./client")))
	log.Fatal(http.ListenAndServe(":8085", nil))
}
