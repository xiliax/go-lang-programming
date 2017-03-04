package main

import (
	"fmt"
)

func main() {
	fmt.Println("Using Channel Selction")

	c := greetings()

	for v := range c {
		fmt.Print(v)
	}
	fmt.Println()
}

func greetings() <-chan string {
	out := make(chan string, 100)

	for i := 0; i < 26; i++ {
		select {
		case out <- "a":
		case out <- "b":
		case out <- "d":
		case out <- "e":
		case out <- "c":
		case out <- "f":
		case out <- "0":
		case out <- "1":
		case out <- "2":
		case out <- "3":
		case out <- "4":
		case out <- "6":
		case out <- "5":
		case out <- "7":
		case out <- "8":
		case out <- "9":
		}
	}

	close(out)
	return out
}
