package p

import "fmt"

func (r Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v, ssn: %v)",
		r.Name, r.Age, r.Ssn)
}
