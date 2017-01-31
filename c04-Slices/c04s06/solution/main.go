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
	count := len(a)
	// TODO - 1 : Print sum of number
	sum := 0
	for _, v := range a {
		sum = sum + v
	}

	fmt.Printf("Sum: %v\n", sum)

	// TODO - 2 : Print average of number
	fmt.Printf("Count of number(s): %v\n", count)

	// TODO - 3 : Print smallest number
	min := a[0]
	for i := 1; i < count; i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	fmt.Printf("Min: %v\n", min)

	// TODO - 4 : Print the lastest number
	max := a[0]
	for i := 1; i < count; i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Printf("Max: %v\n\n", max)

}
