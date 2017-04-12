package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Namer interface {
	SetName(v string)
}

func main() {
	fmt.Println("Implementing Interfaces")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"Jane Doe", 91}

	fmt.Println(bob)
	fmt.Println(jane.String())

	jane.SetName("Mr. Rose") // can i take the address of 'jane', if yes, then (&jane).SetName()
	fmt.Println(jane)

	var pPeter *Person
	pPeter = new(Person)
	pPeter.SetName("Mr. Peters")
	fmt.Println(pPeter) // fmt.Println((*pPeter).String())

	m0 := map[string]*Person{"jane": &jane}
	fmt.Println(m0)

	m0["jane"].SetName("Who's there")

	fmt.Println(m0)
}

func (p *Person) SetName(v string) {
	p.name = v
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v)", p.name, p.age)
}
