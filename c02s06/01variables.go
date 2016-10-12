package main

import "fmt"

func main() {
	var foo = func(s string) {
		fmt.Println("My anonymous funciton prints 's' = ", s)
	}

	foo("Hello, World")
}
