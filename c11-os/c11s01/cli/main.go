package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Creating File with os.Create()")

	f, err := os.Create("data.txt")
	if nil != err {
		log.Fatalln(err)
	}

	io.WriteString(f, "Hello, World!\n")
	io.WriteString(f, "It is a wonderful day!\n")
	io.WriteString(f, "I love Go Language!\n")

	f.Close()
}
