package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Standard In - os.Stdin")

	name := myReadInput("Please enter your name")
	age := myReadInput("Please enter your age")
	fmt.Printf("Hi %v, you are %v years old\n", name, age)
}

func myReadInput(p string) (d string) {
	buf := make([]byte, 100)
	fmt.Printf("%v: ", p)

	n, err := os.Stdin.Read(buf)
	if nil != err {
		return ""
	}

	d = string(buf[:n-1])
	return
}
