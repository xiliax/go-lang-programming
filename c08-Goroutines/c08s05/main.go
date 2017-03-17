package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	c0 := producer("foo")
	c1 := producer("hoo")
	c2 := producer("goo")
	c3 := producer("boo")

	result := consumer(c0, c1, c2, c3)

	go func() {
		for {
			fmt.Println(<-result)
		}
	}()

	time.Sleep(5 * time.Second)
}

func producer(s string) chan string {
	out := make(chan string)

	go func() {

		fmt.Println("Hello, I am the Producer for ", s)
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Nanosecond)
			out <- s
		}
	}()

	return out
}

func consumer(in0, in1, in2, in3 <-chan string) chan string {
	fmt.Println("Hello, I am the Consumer")
	out := make(chan string)
	reader := func(in <-chan string) {
		for v := range in {
			out <- v
		}
	}

	go reader(in0)
	go reader(in1)
	go reader(in2)
	go reader(in3)

	return out
}
