package main

import "fmt"

func main() {
	fmt.Println("Copying Slices with copy()")
	var s = []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'} // hello, world

	fmt.Println(string(s))

	// TODO - 1: assign the slice of bytes 'hello' to a variable

	// TODO - 2: assign the slice of bytes 'world' to another variable

	// TODO - 3: use the copy function to copy first slice to the second

	// ------------ DONE ------------
	// DO NOT CHANGE BELOW THIS LINE
	fmt.Println(string(s))
}
