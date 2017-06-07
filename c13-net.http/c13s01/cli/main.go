package main

import "fmt"
import "net/http"
import "log"
import "io"

var htmlDoc = `
	<html>
		<head>
			<title>Hello from Go net/http</title>
		</head>
		<body>
			<h1>Hello, there</h1>
			<p style="background-color:grey">
			  Thanks for coming by.
			</p>
		</body>
	</html>
`

func main() {
	fmt.Println("Package net/http - A Simple Web Server")

	http.HandleFunc("/", indexHandler)
	log.Fatalln(http.ListenAndServe(":1234", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, htmlDoc)
}
