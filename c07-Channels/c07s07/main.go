package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("More Channel Selction")

	q := "query"

	now := time.Now()
	timeout := time.After(200 * time.Millisecond)

	r0 := search(q)
	r1 := search(q)
	r2 := search(q)

	select {
	case t := <-r0:
		fmt.Printf("Search0 completed after %v\n", t.Sub(now))
	case t := <-r1:
		fmt.Printf("Search1 completed after %v\n", t.Sub(now))
	case t := <-r2:
		fmt.Printf("Search2 completed after %v\n", t.Sub(now))
	case <-timeout:
		fmt.Printf("Too slow, search timedout after 200ms\n")
	}
}

func search(q string) <-chan time.Time {
	out := time.After(time.Duration(rand.Intn(500)) * time.Millisecond)

	return out
}
