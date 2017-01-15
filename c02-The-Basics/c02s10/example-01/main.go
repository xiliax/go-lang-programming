package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := shortHandVariableDeclaration(useLocalVar(5))
	fmt.Println(s)
	ptr := addressOfLocalVar()
	fmt.Printf("Address of pi: %v\n", ptr)
}

func useLocalVar(i int) string {
	var x = 3
	return "Add function var x to input i: " + strconv.Itoa(i+x)
}

func shortHandVariableDeclaration(s string) string {
	result := "*** " + s + " ***"
	return result
}

func addressOfLocalVar() *float64 {
	pi := 3.145
	return &pi
}

func foo(a, b int) (x, y, z int) {
	x = a
	y = b
	z = x + y
	return
}
