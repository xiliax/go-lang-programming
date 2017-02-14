package main

import "fmt"

// Person defines fields for an individual, such as name, age, and snn, etc.
type Person struct {
	name string
	age  int
	snn  string
	// TODO - add additional fields for a person's address
}

func main() {
	// TODO - create a slice of Person containing 3 persons

	// print each person's name and city/state ONLY
	for i, p := range people {
		fmt.Printf("%v. %v lives in %v/%v\n",
			i, p.name, p.city, p.state)
	}
}
