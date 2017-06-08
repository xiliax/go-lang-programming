package main

import "fmt"
import "net/http"
import "log"
import "io"

func main() {
	fmt.Println("Package net/http - A Simple Web Server")

	http.HandleFunc("/", welcomePageHandler)
	http.HandleFunc("/home", homePageHandler)
	log.Fatalln(http.ListenAndServe(":1234", nil))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, homePage)
}

func welcomePageHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, welcomePage)
}
