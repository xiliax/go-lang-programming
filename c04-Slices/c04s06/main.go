package main

import "fmt"

func main() {
	fmt.Println("Understanding the 'spread operator', ...")
	var s0 = []int{11, 56, 37, 89, 23, 101}

	printNumbers(s0...) // printNumbers(s0[0], s0[1], ..., s0[len(s0) - 1])
	fmt.Println("Second call of printNumbers()")
	printNumbers(11, 56, 37, 89, 23, 101)
}

func printNumbers(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
