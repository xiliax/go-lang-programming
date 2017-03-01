package main

import "fmt"

func main() {
	fmt.Println("Introduction to using Channel")

	var c chan int
	fmt.Printf("c: %v\n", c)

	// check the length
	fmt.Println("len of c: ", len(c))

	c = make(chan int, 2)
	// check the length
	fmt.Println("len of c: ", len(c))

	// reading from a channel
	c <- 10
	c <- 25
	fmt.Println("len of c: ", len(c))

	var v = <-c
	fmt.Println("len of c: ", len(c))
	fmt.Printf("v = %v\n", v)
}
