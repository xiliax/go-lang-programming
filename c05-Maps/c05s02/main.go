package main

import (
	"fmt"
)

func main() {
	fmt.Println("Creating and Using Maps")
	var m map[int]int

	fmt.Printf("m: %v, len: %v\n", m, len(m))
	v := m[5]
	fmt.Printf("m[5]: %v\n", v)
	m = make(map[int]int)
	m[0] = 21
	m[1] = 53
	m[2] = 73
	m[3] = 4
	m[5] = 126
	m[6] = 8
	fmt.Printf("DEBUG: Before Delete: m = %v, len = %v\n", m, len(m))
	delete(m, 2)
	fmt.Printf("DEBUG: After Delete: m = %v, len = %v\n", m, len(m))

	//	fmt.Println(m1)

}
