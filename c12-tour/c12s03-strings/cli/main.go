package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person struct {
	name   string
	age    int
	height float64
}

func main() {
	fmt.Println("Using Package stirngs")

	f, err := os.Open("data.txt")
	if nil != err {
		log.Fatalf("Unable to open file for eading: %v\n", err)
	}
	defer f.Close()

	header, err := getHeader(f)
	if nil != err {
		log.Fatal(err)
	}

	fmt.Printf("%-15s %-4s %-8s\n", header[0], header[1], header[2])

	var p *Person
	for nil == err {
		p, err = getRow(f)
		if nil != p {
			fmt.Printf("%-15s %-4d %-8.1f\n", p.name, p.age, p.height)
		}
	}
}

func getRow(r io.Reader) (p *Person, err error) {
	var n int
	var name string
	var age int
	var height float64

	n, err = fmt.Fscanln(r, &name, &age, &height)
	if n != 3 {
		return nil, err
	}

	if nil != err && err != io.EOF {
		return nil, err
	}
	p = &Person{name: name, age: age, height: height}
	return
}

func getHeader(r io.Reader) (header []string, err error) {
	var input string
	n, err := fmt.Fscanf(r, "%s\n", &input)
	fmt.Println(n, input)
	if nil != err {
		return
	}

	header = strings.Split(input, ",")
	return
}
