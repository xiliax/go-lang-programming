package main

import "fmt"

func main() {
	fmt.Println("Copying Slices using copy()")

	var dst []int
	copy(dst, []int{7, 91})
	printSlice("dst", dst)

	var s = []int{7, 91}
	dst = make([]int, 1)
	copy(dst, s)
	printSlice("dst", dst)

	dst = make([]int, 3)
	copy(dst, s)
	printSlice("dst", dst)

	s = []int{7, 91, 3, 59, 72}
	printSlice("s", s)

	var src = s[:4]
	dst = s[1:4]
	printSlice("src", src)
	printSlice("dst", dst)

	copy(dst, src)
	printSlice("src", src)
	printSlice("dst", dst)
	printSlice("s", s)
}

func printSlice(n string, a []int) {
	fmt.Printf("%v = nums[:] => %#v, len: %v, cap: %v\n", n, a, len(a), cap(a))
}
