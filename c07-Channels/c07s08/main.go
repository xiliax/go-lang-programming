package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ch. 7 - Channels - Review")
	c := time.After(1 * time.Nanosecond)

	for i := 0; i < 140; i++ {

		select {
		case v := <-c:
			fmt.Println("Odd", v)
			break

		default:
			fmt.Println("No more values")
		}
	}
}
