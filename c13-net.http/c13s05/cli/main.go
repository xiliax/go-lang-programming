package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HomePage struct {
	FavoriteItems []string
}

type Photo struct {
	Url   string
	Date  string
	Notes []string
}

type PhotosPage struct {
	Photos []Photo
}

func main() {
	fmt.Println("File Server - Part 2")

	http.HandleFunc("/", welcomePageHandler)
	http.HandleFunc("/home", homePageHandler)
	http.HandleFunc("/photos", photosPageHandler)
	staticHandler := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", staticHandler))
	log.Fatalln(http.ListenAndServe(":1234", nil))
}

func photosPageHandler(w http.ResponseWriter, r *http.Request) {
	t := tmpls.Lookup("photo-wall.html")
	data, err := ioutil.ReadFile("photos.json")
	if nil != err {
		log.Fatalln(err)
	}
	page := PhotosPage{}
	err = json.Unmarshal(data, &page)
	if nil != err {
		log.Fatalln(err)
	}

	t.Execute(w, &page)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	t := tmpls.Lookup("home.html")
	data := &HomePage{
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
