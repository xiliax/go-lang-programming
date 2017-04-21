package main

import (
	"errors"
	"fmt"
)

const SSN_LEN = 11

type Person struct {
	name string
	age  uint8  // 0 - 127
	ssn  string // 999-99-9999  ( 11 byte)
}

func main() {
	fmt.Println("Person Implementation of io.Writer")

	/* Format: [
	   <name-len>:byte,
	   <bytes of name>:[]byte,
	   <age>:byte,
	   <bytes of ssn>:[11]byte
	]
	*/
	var data = []byte{
		5, 66, 111, 98, 97, 121, 35, 57, 57, 49, 45, 57, 48, 45, 49, 48, 51, 50,
	}

	fmt.Println(data)

	var p Person
	fmt.Println(p)

	_, err := p.Write(data)
	fmt.Printf("err: %v, %v\n", err, p)
}

func (r *Person) Write(p []byte) (n int, err error) {
	if r == nil {
		return 0, errors.New("Invalid reciever")
	}

	if 0 == len(p) {
		return 0, errors.New("Invalid parameter")
	}

	nameLen := int(p[0])
	if (nameLen + 1) > len(p) {
		return 0, errors.New("Insufficient data for name")
	}
	r.name = string(p[1 : nameLen+1])

	if (nameLen + 2) > len(p) {
		return 0, errors.New("Insufficient data for age")
	}

	r.age = uint8(p[nameLen+1])

	buf := p[nameLen+2:]
	if SSN_LEN > len(buf) {
		return 0, errors.New("Insufficient data for ssn")
	}

	r.ssn = string(buf[:SSN_LEN])

	return (SSN_LEN + 2 + nameLen), nil
}

func (p Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v, ssn: %v)",
		p.name, p.age, p.ssn)
}
