package main

import "fmt"
import "time"

func main() {
	c0 := producer()
	go consumer(c)
	time.Sleep(1 * time.Millisecond)
}

func producer() chan int {
	out := make(chan int)

	go func() {

		fmt.Println("Hello, I am the Producer")
		for i := 0; i < 20; i++ {
			out <- i * 2
		}
		close(out)

	}()

	return out
}

func consumer(in <-chan int) {
	fmt.Println("Hello, I am the Consumer")
	for {
		select {
		case v := <-in:
			fmt.Println(v)
		default:
			fmt.Println("Waiting for value from generator")
		}
	}
}
