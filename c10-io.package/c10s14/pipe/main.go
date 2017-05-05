package main

import (
	"fmt"
	"io"
	"sync"
)

func main() {
	fmt.Println("Example of io.Pipe()")

	pr, pw := io.Pipe()

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		readStuff(pr)
		pr.Close()
		wg.Done()
	}()

	io.WriteString(pw, "Hello, World")
	pw.Close()

	wg.Wait()
}

func readStuff(r io.Reader) {
	buf := make([]byte, 100)
	r.Read(buf)
	fmt.Printf("Read in readStuff: %v\n", buf)
}
