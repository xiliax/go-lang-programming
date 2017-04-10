package main

import "fmt"

type Person struct {
	name string
	age  int
}

type OtherStringer interface {
	String() string
}

type Stringer interface {
	String() string
}

func main() {
	fmt.Println("Implementing Interfaces")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"Jane Doe", 91}

	fmt.Println(bob)
	fmt.Println(jane)

	useMainOtherStringer(bob)
	useMainStringer(jane)

}

func useMainOtherStringer(s OtherStringer) {
	fmt.Printf("This object implements main.OtherStringer interface - > %v\n", s)
}

func useMainStringer(s Stringer) {
	fmt.Printf("This object implements main.Stringer interface - > %v\n", s)
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v)", p.name, p.age)
}
