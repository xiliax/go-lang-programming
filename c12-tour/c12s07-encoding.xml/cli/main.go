package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name   string
	Age    int
	Height float64
}

var (
	inDataFile  = "data.csv"
	outDataFile = "data.xml"
)

// encode Person object to XML
func main() {
	fmt.Println("Using encoding/xml standard package")

	personCh := readCSVData(inDataFile)

	done, err := writeXMLData(outDataFile, personCh)
	if nil != err {
		log.Fatalln(err)
	}

	<-done
	fmt.Printf("Finish reading data from %s to %s successfully\n", inDataFile, outDataFile)
}
func writeXMLData(filename string, in <-chan *Person) (done <-chan bool, err error) {
	f, err := os.Create(outDataFile)
	if nil != err {
		return
	}

	c := make(chan bool)
	go func() {
		enc := xml.NewEncoder(f)
		enc.Indent("  ", "  ")
		for p := range in {
			// encode p as JSON and write to f
			enc.Encode(p)
		}
		f.Close()
		c <- true // we are done writing output
		close(c)
	}()

	return c, nil
}
func readCSVData(filename string) (ch <-chan *Person) {
	linesCh := readFile(filename)
	header := getHeader(linesCh)
	personCh := getRow(linesCh)

	fmt.Printf("%-15s %4s %8s\n",
		fixHeader(header[0]), fixHeader(header[1]),
		fixHeader(header[2]))

	c := make(chan *Person)
	// copy data to output and to another channel
	go func() {
		for p := range personCh {
			fmt.Printf("%-15v %4v %8.1f\n", p.Name, p.Age, p.Height)
			c <- p
		}
		close(c)
	}()
	return c
}

func fixHeader(s string) string {
	s = strings.ToLower(s)
	return strings.Title(s)
}

// read lines from channel, produce *Person on out channel
func getRow(in <-chan string) (ch <-chan *Person) {
	c := make(chan *Person)
	go func() {
		var name string
		var age int
		var height float64
		var p *Person

		for line := range in {
			fields := strings.Split(line, ",")
			p = &Person{}
			switch len(fields) {
			case 3:
				fmt.Sscan(fields[2], &height)
				p.Height = height
				fallthrough
			case 2:
				fmt.Sscan(fields[1], &age)
				p.Age = age
				fallthrough
			case 1:
				name = strings.TrimSpace(fields[0])
				p.Name = strings.Title(strings.ToLower(name))
			}

			c <- p
		}
		close(c)
	}()

	return c
}

// this is a blocking function, it reads one line from the channel and returns
func getHeader(in <-chan string) (header []string) {
	input := <-in // read one line
	input = strings.TrimSpace(input)
	header = strings.Split(input, ",")
	return
}

// read a line from a file until EOF or error and send it
// on a read-only channel
func readFile(filename string) (ch <-chan string) {
	c := make(chan string)

	// generate the lines for the channel
	go func() {
		f, err := os.Open(filename)
		if nil != err {
			return // ch is nil
		}

		bio := bufio.NewReader(f)

		var input string
		for nil == err {
			input, err = bio.ReadString('\n')
			// exit loop if any error (regardless of type)
			// but we might still have read some data, so send it
			c <- input // have to use 'c(r/w)' and not 'ch(ro)'
		}

		// close channel for after we finished reading from file
		close(c)
		f.Close()
	}()

	ch = c
	return
}
