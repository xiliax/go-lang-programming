package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Printer interface {
	Print()
}

type Duck interface {
	Waddle()
	Quack()
}

type IndianRunner struct {
	name string
}

func main() {
	fmt.Println("Empty Interface")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"Jane Doe", 91}
	m0 := map[string]Person{"first": jane}

	myPrinter(5)
	myPrinter("Hello, World")
	myPrinter(m0)
	myPrinter(bob)
	myPrinter(jane)
	myPrinter(IndianRunner{name: "quick"})
}

func myPrinter(o interface{}) {
	fmt.Printf("MyPrinter: - %T -> %v\n", o, o)
}

func (i IndianRunner) String() string {
	return fmt.Sprintf("Duck (name: %v)", i.name)
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v)", p.name, p.age)
}
