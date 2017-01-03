package main

import "fmt"

func main() {
	fmt.Println("Growing Slices at Runtime with 'append()'")
	var s0 []int
	printArray("s0", s0)
	s0 = make([]int, 2, 3)

	printArray("s0", s0)
	s0[0] = 17
	s0[1] = 91
	printArray("s0", s0)

	s1 := myAppend(s0, 5)
	printArray("s1", s1)

	s2 := myAppend(s0, 5, 42, 81, 63)
	printArray("s0", s0)
	printArray("s2", s2)
}

func printArray(n string, a []int) {
	fmt.Printf("%v => %#v, len: %v, cap: %v\n", n, a, len(a), cap(a))
}

func myAppend(s []int, e ...int) []int {
	var t []int
	if len(e) > (cap(s) - len(s)) {
		t = make([]int, cap(s)+len(e))
		for i := 0; i < len(s); i++ {
			t[i] = s[i]
		}
	} else {
		t = s[:len(s)+len(e)]
	}

	for i := 0; i < len(e); i++ {
		t[len(s)+i] = e[i]
	}

	return t
}
