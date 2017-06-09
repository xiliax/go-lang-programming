package main

import (
	"fmt"
	"log"
	"net/http"
)

type Page struct {
	FavoriteItems []string
}

func main() {
	fmt.Println("A Simple Web Server and http/template")

	http.HandleFunc("/", welcomePageHandler)
	http.HandleFunc("/home", homePageHandler)
	log.Fatalln(http.ListenAndServe(":1234", nil))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t := tmpls.Lookup("home.html")
	data := &Page{
		[]string{
			"Reading",
			"Programming in Go Language",
			"Watching Movies",
			"Hiking",
			"Riding",
		},
	}
	t.Execute(w, data)
}

func welcomePageHandler(w http.ResponseWriter, r *http.Request) {
	t := tmpls.Lookup("welcome.html")
	t.Execute(w, nil)
}
