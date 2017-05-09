package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Opening and Reading Files")

	f, err := os.Open("data.txt")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	buf := make([]byte, 200)
	n, err := f.Read(buf)
	if nil != err {
		log.Fatalln(err)
	}

	fmt.Println("Bytes read:", n)
	fmt.Println(string(buf))
}
