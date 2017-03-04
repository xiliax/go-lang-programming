package main

import "fmt"

func main() {
	fmt.Println("Functions Returning Channels")

	c := producer()
	c = processor1(c)
	c = processor2(c)

	for v := range c {
		fmt.Printf("\tv = %v\n", v)
	}
}

func processor2(in <-chan int) <-chan int {
	out := make(chan int, 10)
	for v := range in {
		out <- (v * -5) + 100
	}

	close(out)
	return out
}

func processor1(in <-chan int) <-chan int {
	out := make(chan int, 10)
	for v := range in {
		out <- (v * 2) - 7
	}

	close(out)
	return out
}

func producer() <-chan int {
	out := make(chan int, 10)
	out <- 0
	out <- 1
	out <- 2
	out <- 4
	out <- 8
	out <- 16
	close(out)
	return out
}
