package main

import "fmt"

// Person defines fields for an individual, such as name, age, and snn
type Person struct {
	name string
	age  int
	snn  string
	Address
}

// Address defines the fields for an address
type Address struct {
	street1 string
	street2 string
	city    string
	state   string
	Zip
}

// Zip defines a Zip code (eg: 12345-1105)
type Zip struct {
	code  string
	pobox string
}

func main() {
	// TODO: 1 - declare the map variable as 'people', string -> Person

	// TODO: 2 - allocate/make the map

	// TODO: 3 - Store 3 persons in the map 'people', keyed on the peron's name
	var p Person
	p = Person{"Mary", 35, "011-13-4567",
		Address{"1 First St", "", "Brooklyn", "NY",
			Zip{"11212", ""}},
	}
	// TODO: 4 - Store person 1

	p = Person{"Sam", 18, "011-32-9819",
		Address{"773 George Rd", "", "Queens", "NY",
			Zip{"11235", ""}},
	}
	// TODO: 5 - Store person 1

	p = Person{"Greg", 39, "011-34-4113",
		Address{"93 Kings Ct", "Suite A", "Fargo", "MN",
			Zip{"29019", ""}},
	}
	// TODO: 6 - Store person 1

	for _, p = range people {
		fmt.Printf("%v lives in %v, %v %v\n", p.name, p.city, p.state, p.code)
	}
}
