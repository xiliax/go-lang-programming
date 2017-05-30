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
	fmt.Println("Using strings.Split,Title")

	f, err := os.Open("data.txt")
	if nil != err {
		log.Fatalf("Unable to open file for eading: %v\n", err)
	}
	defer f.Close()

	header, err := getHeader(f)
	if nil != err {
		log.Fatalln(err)
	}

	fmt.Printf("%-15s %4s %8s\n",
		strings.Title(header[0]), strings.Title(header[1]),
		strings.Title(header[2]))

	var p *Person
	for nil == err {
		p, err = getRow(f)

		if err == io.EOF {
			return
		}

		if nil != err {
			log.Fatalf("Error reading row: %v\n", err)
		}

		fmt.Printf("%-15v %4v %8.1f\n", p.name, p.age, p.height)
	}
}
func getRow(r io.Reader) (p *Person, err error) {
	var name string
	var age int
	var height float64

	_, err = fmt.Fscanln(r, &name, &age, &height)

	if nil != err {
		return
	}

	p = &Person{name: name, age: age, height: height}
	return
}

func getHeader(r io.Reader) (header []string, err error) {
	var input string
	_, err = fmt.Fscanln(r, &input)
	if nil != err {
		return
	}

	header = strings.Split(input, ",")
	return
}
