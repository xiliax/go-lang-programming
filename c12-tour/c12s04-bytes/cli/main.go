package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Using 'bytes' standard package")

	var b bytes.Buffer
	(&b).Write([]byte("Hello, World!\n"))
	fmt.Fprintf(&b, "Holy %s\n", "smokes batman!")
	fmt.Println(b.String())

	f, err := os.Open("main.go")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	b.ReadFrom(f)
	fmt.Fprintln(&b, "This is the end!!!")

	lines := bytes.Split(b.Bytes(), []byte("\n"))
	for n, l := range lines {
		fmt.Printf("%-4d %s\n", n+1, string(l))
	}

}
