package main

import "fmt"

// Person defines fields for an individual
type Person struct {
	name string
	age  int
	ssn  string
	dob  string
}

func main() {
	fmt.Println("Introduction to Structs")

	var p Person

	fmt.Printf("%#v\n", p)

	p.name = "Smith"
	p.age = 53
	p.ssn = "111-01-1234"
	fmt.Printf("%v\n", p)

	var p1 = Person{
		age:  21,
		name: "Joan",
		ssn:  "011-21-5431",
	}

	fmt.Printf("%v\n", p1)

	var friends []Person

	friends = append(friends, Person{
		age:  56,
		name: "Mary",
		ssn:  "212-43-8761",
	})

	fmt.Printf("%v\n", friends)

	var people = []Person{
		{
			age:  18,
			name: "Jane",
			ssn:  "013-35-0153",
		},
		{
			age:  29,
			name: "Greg",
			ssn:  "071-43-2053",
		}, {
			age:  40,
			name: "Sam",
			ssn:  "091-43-8761",
		},
	}

	fmt.Printf("%v\n", people)

}
