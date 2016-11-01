package main

import "fmt"

// package level variables
var allItems = 0
var isLoggingEnabled bool = true
var currentUer string = "Verrol"
var chineseChar rune = '世' // "world"

func main() {
	fmt.Println("Package variables:")
	fmt.Println("---------------------------")
	fmt.Println(allItems)
	fmt.Println(isLoggingEnabled)
	fmt.Println(currentUer)
	fmt.Println(chineseChar)

	explicitTypedVariables()
	inferredTypeVariables()
}

func explicitTypedVariables() {
	fmt.Println("Explicitly typed variables:")
	fmt.Println("---------------------------")
	var printingEnabled bool = true
	var itemCount int = 537
	var price float32 = 235.87
	var filename string = "foo.bar.txt"

	printTypeAndValue(printingEnabled)
	printTypeAndValue(itemCount)
	printTypeAndValue(price)
	printTypeAndValue(filename)
}

func inferredTypeVariables() {
	fmt.Println("\nInferred type variables:")
	fmt.Println("---------------------------")
	var printingEnabled = true
	var itemCount = 537
	var price = 235.87
	var filename = "foo.bar.txt"

	printTypeAndValue(printingEnabled)
	printTypeAndValue(itemCount)
	printTypeAndValue(price)
	printTypeAndValue(filename)
}

func printTypeAndValue(t interface{}) {
	fmt.Printf("value: %v, type: %T\n", t, t)
}
