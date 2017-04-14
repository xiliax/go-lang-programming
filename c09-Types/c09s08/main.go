package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Duck interface {
	Waddle()
	Quack()
}

type IndianRunner struct {
	name string
}

func main() {
	fmt.Println("Type Assertions")

	var bob = Person{"Bob Jones", 35}
	var speedy = IndianRunner{"Speedy"}

	useDuck(bob)
	useDuck(speedy)
	useDuck(5)

}

func useDuck(o interface{}) {
	// fmt.Printf("Using Duck: - %T -> %v\n", o, o)

	switch o.(type) {
	case int:
		fmt.Println("I have an 'int' value in 'o'", o)
	case IndianRunner:
		fmt.Println("I have an IndianRunner as the type of the value in 'o'", o)
	case Person:
		fmt.Println("I have an Person as the type of the value in 'o'", o)
	default:
		fmt.Println("Unsupported type")
	}
}

func (p Person) Quack() {
	fmt.Printf("%v quacking", p)
}

func (p Person) Waddle() {
	fmt.Printf("%v waddling", p)
}

func (i IndianRunner) Quack() {
	fmt.Printf("%v quacking", i)
}

func (i IndianRunner) Waddle() {
	fmt.Printf("%v waddling", i)
}

func (i IndianRunner) String() string {
	return fmt.Sprintf("Duck (name: %v)", i.name)
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v)", p.name, p.age)
}
