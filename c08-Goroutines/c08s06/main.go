package main

import "fmt"

const N = 200000

func main() {
	c := make(chan int)
	in := c
	var out chan int

	for i := 0; i < N; i++ {
		out = worker(in)
		in = out
	}

	c <- 0
	fmt.Println(<-out)
}

func worker(in chan int) chan int {
	out := make(chan int)

	go func() {
		v := <-in
		out <- 1 + v
	}()

	return out
}
