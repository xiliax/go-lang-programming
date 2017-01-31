package main

import "fmt"

func main() {
	fmt.Println("Creating Slices at Runtime")
	var s0 []int
	fmt.Printf("s0 = nums[:] => %#v, len: %v, cap: %v\n", s0, len(s0), cap(s0))
	s0 = make([]int, 5)
	fmt.Printf("s0 = nums[:] => %#v, len: %v, cap: %v\n", s0, len(s0), cap(s0))
	s0 = make([]int, 5, 8)
	fmt.Printf("s0 = nums[:] => %#v, len: %v, cap: %v\n", s0, len(s0), cap(s0))
	s1 := s0[:cap(s0)]
	fmt.Printf("s1 = nums[:] => %#v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	s1[5] = 17
	s1[3] = 91
	printArray(s0)
	printArray(s1)
	var testScores []int
	// prompt the user for the nubmer test scores
	var count = 20
	testScores = make([]int, count)
	testScores[0] = 12
	testScores[1] = 56
	printArray(testScores)
}

func printArray(a []int) {
	fmt.Printf("a = nums[:] => %#v, len: %v, cap: %v\n", a, len(a), cap(a))

}
