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
	name    string
	street1 string
	street2 string
	city    string
	state   string
	Zip
}

// Zip is a zip code
type Zip struct { // eg: 12345-1111
	name  string
	code  string
	pobox string
}

func main() {
	var people []Person
	var p Person

	p = Person{"Mary", 35, "011-13-4567",
		Address{"Address1", "1 First St", "", "Brooklyn", "NY",
			Zip{"Zip1", "11212", "1015"}},
	}
	people = append(people, p)

	p = Person{"Sam", 18, "011-32-9819",
		Address{"Address2", "773 George Rd", "", "Queens", "NY",
			Zip{"Zip2", "11235", "2019"}},
	}
	people = append(people, p)

	p = Person{"Greg", 39, "011-34-4113",
		Address{"Address3", "93 Kings Ct", "Suite A", "Fargo", "MN",
			Zip{"Zip3", "29019", ""}},
	}
	people = append(people, p)

	for _, p = range people {
		fmt.Printf("%v lives in %v: [%v, %v %v], Zip name: %v\n", p.name, p.Address.name, p.city, p.state,
			p.code, p.Zip.name)
	}
}
