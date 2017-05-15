package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Using os.Args")

	for _, a := range os.Args[1:] {
		t := getFileType(a)
		fmt.Printf("%v = %v\n", a, t)
	}
}

func getFileType(f string) (t string) {
	fi, err := os.Stat(f)
	if nil != err {
		return "file or directory doesn't exisit"
	}

	if fi.IsDir() {
		return "is a directory"
	}

	return "is a file"
}
