// Demonstrate Array iteration
package main

import (
	"fmt"
	"log"
)

const size = 10

var testResults = [size]int{12, 53, 86, 94, 75, 103, 0, 0, 0, 408}

func main() {
	log.Println("Iteration using 'for' and range")

	for i := 0; i < len(testResults); i++ {
		fmt.Print(testResults[i], " ")
	}
	fmt.Println()

	for i, _ := range testResults {
		fmt.Println(i, "=>")
	}
}
