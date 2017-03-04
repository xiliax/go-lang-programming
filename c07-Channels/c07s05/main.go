package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("More Channel Stuff")

	start := time.Now()
	fmt.Printf("Start: %v\n", start)

	c := time.After(2 * time.Second)

	fmt.Printf("\nEnd: %v\n", <-c)

}
