package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	fmt.Println(6)
	fmt.Println("The value 6:", 6)
	fmt.Println("3 + 1 =", 3+1)
	fmt.Println("2016 divided by 9 divided by 15 =", 2016/9/15)
	fmt.Println("9 divided by 15 = ", 9/15)
	fmt.Println("9.0 divided by 15 =", 9.0/15)
	fmt.Println("The value of Pi =", 3.141592)
	fmt.Println("The statement 'pigs can fly' is", false)
	fmt.Println("My comlex number is", 11+4i)
	fmt.Println("An slice of 4 integers:", []int{1, 2, 3, 5})
	fmt.Println("A 3 element array of string, with only 2 elements initialized:", [3]string{"hello", "world"})
}
