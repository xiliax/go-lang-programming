package main

import "fmt"

func main() {
	fmt.Println("Using Receive-only Channels")

	c := make(chan int, 10)

	producer(c)
	consumer(c)

}

func consumer(c <-chan int) {
	fmt.Println("Reading from a channel")
	for v := range c {
		fmt.Printf("\tv = %v\n", v)
	}
}

func producer(c chan<- int) {
	fmt.Println("Writing to a channel")
	c <- 323
	c <- 10
	c <- 25
	close(c)
}
