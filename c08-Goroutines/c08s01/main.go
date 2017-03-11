package main

import "fmt"

func main() {
	fmt.Println("Introducting 'go' Statements")

	c := make(chan int)

	go producer(c)
	consumer(c)

}

func consumer(c chan int) {
	for v := range c {
		fmt.Printf("Read: v = %v\n", v)
	}
}

var producer = func(c chan int) {

	fmt.Println("Writing to a channel")
	c <- 2

	fmt.Println("Writing to a channel")
	c <- 5

	fmt.Println("Writing to a channel")
	c <- 1

	fmt.Println("Writing to a channel")
	c <- 90

	close(c)
}
