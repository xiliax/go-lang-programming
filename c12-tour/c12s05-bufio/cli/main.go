package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	name   string
	age    int
	height float64
}

func main() {
	fmt.Println("Using bufio standard package")

	linesCh := readFile("data.csv")

	header := getHeader(linesCh)

	fmt.Printf("%-15s %4s %8s\n",
		fixHeader(header[0]), fixHeader(header[1]),
		fixHeader(header[2]))

	personCh := getRow(linesCh)
	for p := range personCh {
		fmt.Printf("%-15v %4v %8.1f\n", p.name, p.age, p.height)
	}
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

		for line := range in {
			fields := strings.Split(line, ",")

			switch len(fields) {
			case 3:
				fmt.Sscan(fields[2], &height)
				fallthrough
			case 2:
				fmt.Sscan(fields[1], &age)
				fallthrough
			case 1:
				name = strings.TrimSpace(fields[0])
				name = strings.Title(strings.ToLower(name))
			}

			c <- &Person{name: name, age: age, height: height}
			name = ""
			age = 0
			height = 0.0
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
