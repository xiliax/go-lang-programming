package main

import (
	"fmt"
	"log"
	"os"

	"errors"

	"github.com/another/c10-io.package/c10s12/p"
)

func main() {
	fmt.Println("Binary File I/O")

	p0 := p.Person{Name: "Jane Doe", Age: 25, Ssn: "011-11-0123"}

	fn := "person.db"
	n, err := writePerson(fn, p0)
	if nil != err {
		log.Fatalln("Write Person err: ", err)
	}
	fmt.Printf("Wrote %v bytes\n", n)

	p1, n, err := readPerson(fn)
	if nil != err {
		log.Fatalln("Read Person err: ", err)
	}

	fmt.Println(p1)
}

func writePerson(fn string, o p.Person) (n int, err error) {
	var f *os.File
	fi, err := os.Stat(fn)
	if nil != err {
		f, err = os.Create(fn)
	} else {
		// ensure file is a regular file
		if nil != fi && fi.IsDir() {
			return 0, errors.New(fn + " is a directory, need a file")
		}

		f, err = os.OpenFile(fn, os.O_WRONLY, 0644)
		if nil != err {
			return
		}
	}

	defer f.Close()

	buf := make([]byte, 100)
	n, err = o.Read(buf)
	if nil != err {
		return
	}
	n, err = f.Write(buf)
	return
}
func readPerson(fn string) (o *p.Person, n int, err error) {
	f, err := os.Open(fn)
	if nil != err {
		return
	}

	defer f.Close()

	buf := make([]byte, 200)
	n, err = f.Read(buf)
	if nil != err {
		return
	}
	fmt.Printf("Read %v bytes\n", n)

	o = new(p.Person)
	n, err = o.Write(buf)
	return
}
