package main

import "fmt"

func main() {
	fmt.Println("Creating Slices")
	var nums = [...]int{12, 53, 86, 94, 75, 10, 2, 15, 37, 40}

	var mid = len(nums) / 2
	var s0 = nums[:]    // {12, 53, 86, 94, 75, 10, 2, 15, 37, 40}
	var s1 = nums[:mid] // {12, 53, 86, 94, 75}
	var s2 = nums[mid:] // {0, 2, 15, 37, 40}

	fmt.Printf("nums => %#v\n", nums)
	fmt.Printf("s0 = nums[:] => %#v\n", s0)
	fmt.Printf("s1 = nums[:len(nums)/2] => %#v\n", s1)
	fmt.Printf("s2 = nums[len(nums)/2:] => %#v\n", s2)

	var s3 = nums[3:7]
	fmt.Printf("s3 = nums[3:7] => %#v\n", s3) // {94, 75, 10, 2}

}
