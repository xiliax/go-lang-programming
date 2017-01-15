package main

import "fmt"

func main() {
	foo()
	myPrintString("Hello world!")
	fmt.Println("myAdd of 1, 5 is", myAdd(1, 5))
	var a, b = swap(5, 1)
	fmt.Printf("Swap of 5, 1 is %v, %v\n", a, b)
	printArgs()
	printArgs("Hello")
	printArgs("Hello", "World")
	printArgs("Hello", "World", "!")
}

func foo() {
	fmt.Println("I am the function 'foo()'")
}

func myPrintString(value string) {
	fmt.Printf("-- value: %v ---\n", value)
}

func myAdd(a, b int) int {
	return (a + b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func printArgs(args ...string) {
	fmt.Println("Number of args to function 'printArgs(...)' are:", len(args))
}
