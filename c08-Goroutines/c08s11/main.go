package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of OS Threads: ", runtime.GOMAXPROCS(0))
}
