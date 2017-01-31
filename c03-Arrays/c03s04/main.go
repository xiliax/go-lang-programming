// Demonstrate using arrays to hold a collection values (of the same)
package main

import (
	"fmt"
	"log"
)

const size = 11

func main() {
	log.Println("Using Arrays and Functions")

	var data = getTestData()
	printTestResults(&data)
	fmt.Println("main.data:", data)

}

func printTestResults(testResults *[size]int) {
	testResults[8] = 512
	i := 6
	var s0 = testResults[2:i]
	s0[2] = 1010
	var s1 = testResults[i:]
	s1[2] = 9999

	fmt.Printf("Length Array: %#v\n", len(testResults))
	fmt.Printf("Lower Array: %#v\n", s0)
	fmt.Printf("Upper Array: %#v\n", s1)
	fmt.Printf("Length s0: %#v\n", len(s0))

}

func getTestData() (data [size]int) {

	data = [size]int{12, 53, 86, 94, 75, 103, 0, 0, 0, 408}
	return
}
