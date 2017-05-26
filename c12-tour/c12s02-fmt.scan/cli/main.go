package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Using fmt.ScanX")

	var name string
	var age int
	var height float64

	f, err := os.Open("data.txt")
	if nil != err {
		log.Fatalf("Unable to open file for eading: %v\n", err)
	}
	defer f.Close()

	var n int
	for nil == err {
		n, err = fmt.Fscanln(f, &name, &age, &height)
		if n > 0 {
			fmt.Printf("Hi, %v you are %v years old and %.1f ft/in\n", name, age, height)
		}

		if nil != err && err != io.EOF {
			log.Fatalf("Read %v values, err: %v\n", n, err)
		}
	}
}
