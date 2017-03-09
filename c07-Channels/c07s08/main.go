package main

import (
	"fmt"
	"time"
)

func main() {
	var later = time.After((1 / 10) * time.Nanosecond)

	for i := 0; i < 10; i++ {

		select {
		case <-later:
			fmt.Printf("Timer expired after 1ns\n")

		default:
			fmt.Println("I am the default")
		}
	}
}
