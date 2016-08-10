package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello, world!")
	<-time.After(5 * time.Second)
	fmt.Println("hello again")
}
