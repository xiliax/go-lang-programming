package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(1 * time.Millisecond)
}

func producer(out chan<- int) {
	fmt.Println("Hello, I am the Producer")
	for i := 0; i < 20; i++ {
		out <- i * 2
	}
	//close(out)
}
func consumer(in <-chan int) {
	fmt.Println("Hello, I am the Consumer")
	for {
		select {
		case v := <-in:
			fmt.Println(v)
		default:
			fmt.Println("Waiting for value from Producer")
		}
	}
}
