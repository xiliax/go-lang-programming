package main

import "fmt"

type Person struct {
	name string
	age  int
}

type NameChanger interface {
	ChangeName(v string)
}

func main() {
	fmt.Println("Implementing Interfaces")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"Jane Doe", 91}

	fmt.Println(bob)
	fmt.Println(jane.String())

	jane.ChangeName("Mr. Gold")
	fmt.Println(jane)
}

func (p *Person) ChangeName(v string) {
	p.name = v
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v)", p.name, p.age)
}
