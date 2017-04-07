package main

import "fmt"

type Person struct {
	name string
	age  int
}

type IndianRunner struct {
	name   string
	weight int
}

type Watervale struct {
	name   string
	weight int
}

type Duck interface {
	Quack()
	Waddle()
}

func main() {
	fmt.Println("Interface type")

	var bob = Person{"Bob Jones", 35}
	var jane = Person{"Jane Doe", 91}

	fmt.Println(bob)
	fmt.Println(jane)

	var duck0 = IndianRunner{"Speedy", 7}
	fmt.Println(duck0)
	duck0.Quack()

	var duck1 = Watervale{"Swimmer", 6}
	fmt.Println(duck1)

	useDuck(duck0)
}

func useDuck(d Duck) {
	d.Quack()
	d.Waddle()
}
func (d Watervale) String() string {
	return fmt.Sprintf("My Watervale name is %v and weights %vlbs", d.name, d.weight)
}

func (d IndianRunner) Waddle() {
	fmt.Printf("My Indian Runner %v is waddling\n", d.name)
}

func (d IndianRunner) Quack() {
	fmt.Printf("My Indian Runner %v is quacking\n", d.name)
}

func (d IndianRunner) String() string {
	return fmt.Sprintf("My Indian Runner name is %v and weights %vlbs", d.name, d.weight)
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v)", p.name, p.age)
}
