package main

import (
	"fmt"
	"time"
)

const delay = 200

func main() {
	fmt.Println("Introduction to Goroutines : Walking and Chewing Gum")
	go chewing("gum")
	walking()
}

func walking() {
	fmt.Println("start walking")
	for i := 0; i < 10; i++ {
		fmt.Println("\twalking...")
		time.Sleep(delay * time.Millisecond)
	}
	fmt.Println("finish walking")
}

func chewing(w string) {
	fmt.Println("starting chewing", w)
	for i := 0; i < 10; i++ {
		fmt.Println("\tchewing", w)
		time.Sleep(delay * time.Millisecond)
	}
	fmt.Println("finish chewing", w)
}
