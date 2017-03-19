package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 200000

func main() {
	quit := make(chan bool)
	c0 := chatter("Bob", quit)
	c1 := chatter("Joe", quit)
	c2 := chatter("Anne", quit)
	c3 := chatter("Mary", quit)

	go func() {
		for {
			select {
			case v := <-c0:
				fmt.Println(v)
			case v := <-c1:
				fmt.Println(v)
			case v := <-c2:
				fmt.Println(v)
			case v := <-c3:
				fmt.Println(v)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- true
	quit <- true
	quit <- true
	quit <- true
}

func chatter(s string, quit <-chan bool) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case <-quit:
				// clean up
				fmt.Println(s, " is quitting")
				break
			case out <- s:
				time.Sleep(time.Duration(rand.Intn(100)) * time.Nanosecond)
			}
		}
	}()

	return out
}
