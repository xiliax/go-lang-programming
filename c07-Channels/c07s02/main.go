package main

import "fmt"

func main() {
	fmt.Println("Using Send-only Channels")

	c := make(chan int, 10)

	producer(c)

	// reading from a channel
	for v := range c {
		fmt.Printf("v = %v\n", v)
	}
}

func producer(c chan<- int) {
	// writing to a channel
	c <- 323
	c <- 10
	c <- 25
	close(c)
}
