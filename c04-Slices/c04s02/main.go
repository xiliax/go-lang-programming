package main

import "fmt"

func main() {
	fmt.Println("Iteration Using 'for' and range()")
	var nums = [...]int{12, 53, 86, 94, 75, 10, 2, 15, 37, 40}

	for i := 0; i < len(nums); i++ {
		fmt.Println(i, "=>", nums[i])
	}

	fmt.Println("----------------------------")
	for _, v := range nums {
		fmt.Println("=>", v)
	}
}
