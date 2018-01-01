// Demonstrate using arrays to hold a collection values (of the same type)
package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Introduction to Arrays")
	printTestResults()
}

func printTestResults() {
	// 12, 53, 86, 94, 75, 103
	var testResults [10]int
	testResults[0] = 12
	testResults[1] = 53
	testResults[2] = 86
	testResults[3] = 94
	testResults[4] = 75
	testResults[5] = 103
	i := 10
	testResults[i] = 408

	fmt.Println("Test Results: ", testResults)
}
