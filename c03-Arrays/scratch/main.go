package main

import "fmt"

func main() {
	fmt.Println("Growing Slices at Runtime with 'append()'")
	var s = []int{7, 91}
	var s0 = s[:]
	var d0 = []int{}

	printArray("s0", s0)
	printArray("d0", d0)
	copy(d0, s0)
	printArray("d0", d0)

	d0 = []int{0}

	printArray("s0", s0)
	printArray("d0", d0)
	copy(d0, s0)
	printArray("d0", d0)

	d0 = []int{0, 0}

	printArray("s0", s0)
	printArray("d0", d0)
	copy(d0, s0)
	printArray("d0", d0)

	s = []int{7, 91, 3, 59, 72}

	printArray("s", s)

	var src = s[0:5]
	var dst = s[1:4]

	printArray("src", src)
	printArray("dst", dst)
	copy(dst, src)
	printArray("src", src)
	printArray("dst", dst)
	printArray("s", s)
}

func printArray(n string, a []int) {
	fmt.Printf("%v = nums[:] => %#v, len: %v, cap: %v\n", n, a, len(a), cap(a))
}

func myAppend(s []int, e ...int) []int {
	fmt.Println("s len: ", len(s))
	fmt.Println("s cap: ", cap(s))
	fmt.Println("e len: ", len(e))
	if len(e) > (cap(s) - len(s)) {
		t := make([]int, cap(s)+len(e))
		for i := 0; i < len(s); i++ {
			t[i] = s[i]
		}
		for i := 0; i < len(e); i++ {
			t[len(s)+i] = e[i]
		}
		s = t
	} else {
		t := s[:len(s)+len(e)]
		for i := 0; i < len(e); i++ {
			t[len(s)+i] = e[i]
		}
		s = t
	}

	return s
}
