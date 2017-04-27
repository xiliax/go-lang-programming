package main

import (
	"errors"
	"fmt"
	"log"
)

type Person struct {
	name string
	age  uint8  // no more 255 years
	ssn  string // 11-chars (999-99-9999)
}

func main() {
	fmt.Println("Person Reader")

	bob := Person{"Bobby Jones", 35, "100-01-1101"}
	mary := Person{"Mary-Anne Smith", 83, "831-01-8119"}
	jane := Person{"Jane Doe", 17, "021-23-2376"}

	var buf = make([]byte, 110)
	offset, n := 0, 0
	var err error

	n, err = bob.Read(buf[offset:])
	offset += n

	n, err = mary.Read(buf[offset:])
	offset += n

	n, err = jane.Read(buf[offset:])
	if nil != err {
		fmt.Println(buf)
		log.Fatalf("n: %v, err: %v\n", n, err)
	}
	offset += n

	n, err = mary.Read(buf[offset:])
	if nil != err {
		fmt.Println(buf)
		log.Fatalf("n: %v, err: %v\n", n, err)
	}
	offset += n

	n, err = jane.Read(buf[offset:])
	if nil != err {
		fmt.Println(buf)
		log.Fatalf("n: %v, err: %v\n", n, err)
	}
	offset += n

	fmt.Println(buf)
}

/*
	[x,x,x,x,x,x,x,x,x,......,x,x,x]
	Format: [
				<name-len>: byte
				<name>: []byte
				<age: byte
				<ssn>:[]byte
			]
*/
func (r Person) Read(p []byte) (n int, err error) {
	// do the work of encoding a person 'r' into bytes and store them in 'p'

	var l = len(r.name)
	var buf = make([]byte, l+1)

	// encode r.name
	buf[0] = byte(l)
	for i := 0; i < l; i++ {
		buf[i+1] = r.name[i]
	}

	// encode r.age
	buf = append(buf, byte(r.age))

	// encode r.ssn
	l = len(r.ssn)
	for i := 0; i < l; i++ {
		buf = append(buf, r.ssn[i])
	}

	if len(p) < len(buf) {
		s := fmt.Sprintf("Buffer not big enough, required: %v, provided: %v",
			len(buf), len(p))
		return 0, errors.New(s)
	}
	n = copy(p, buf)

	return n, nil
}

func (r Person) String() string {
	return fmt.Sprintf("Person (name: %v, age: %v, ssn: %v)",
		r.name, r.age, r.ssn)
}
