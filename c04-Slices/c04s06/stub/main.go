package main

import "fmt"

func main() {
	fmt.Println("Understanding the 'spread operator', ...")
	var s0 = []int{11, 56, 37, 89, 23, 101}

	numberStats(11)
	numberStats(11, 56)
	numberStats(11, 89, 23)
	numberStats(s0...)
}

func printNumbers(a ...int) {
	for _, v := range a {
		fmt.Println(v)
	}
}

// numberStats - print some stats on the numbers provided
func numberStats(a ...int) {
	// TODO - 1 : Print sum of number

	// TODO - 2 : Print average of number

	// TODO - 3 : Print smallest number

	// TODO - 4 : Print the lastest number
}
