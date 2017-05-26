package main

import "github.com/another/c12-fmt/c12s01/f"

func main() {
	f.Println("Introduction to fmt")
	name := "Jane"
	age := 35
	f.Printf("%s is %d years old and %.2f feet/inches tall\n", name, age, 5.3)
}
