package main

import "fmt"

// Person defines fields for an individual, such as name, age, and snn, etc.
type Person struct {
	name string
	age  int
	snn  string
	// TODO - add additional fields for a person's address
	street1 string
	street2 string
	city    string
	state   string
	zip     string
}

func main() {
	// TODO - create a slice of Person containing 3 persons
	var people []Person // nil slice
	var p Person
	p = Person{"Mary", 35, "011-13-4567",
		"1 First St", "", "Brooklyn", "NY", "11212",
	}
	people = append(people, p)
	p = Person{"Sam", 18, "011-32-9819",
		"773 George Rd", "", "Queens", "NY", "11235",
	}
	people = append(people, p)
	p = Person{"Greg", 39, "011-34-4113",
		"93 Kings Ct", "Suite A", "Fargo", "MN", "29019",
	}
	people = append(people, p)
	// TODO - print each person's name and city/state ONLY
	for _, p = range people {
		fmt.Printf("%v lives in %v/%v\n", p.name, p.city, p.state)
	}
}
