package main

import (
	"errors"
	"fmt"
)

type Person struct {
	name string
	age  uint8 // 0 - 127
}

type AgeAverager struct {
	sum   uint64
	count uint64
}

func main() {
	fmt.Println("Age Averager Writer")

	bob := Person{"Bobby", 15}
	jane := Person{"Jane Doe", 35}
	mary := Person{"Mary", 7}

	var c AgeAverager

	c.Write(bob.AgeToByte())
	c.Write(jane.AgeToByte())
	c.Write(mary.AgeToByte())
	var p []byte
	n, err := c.Write(p)
	fmt.Println(n, err)

	n, err = c.Write([]byte{210})
	fmt.Println(n, err)

	fmt.Printf("Age Averager: %v\n", c)
}

func (p Person) AgeToByte() []byte {
	var buf = []byte{byte(p.age)}
	return buf
}

func (c *AgeAverager) Write(p []byte) (n int, err error) {
	l := len(p)
	switch {

	case p == nil:
		return 0, errors.New("Invalid parameter")
	case 127 < p[0]:
		return 0, errors.New("Age must be less than or equal to 127")
	}

	c.count++ // c.count = c.count + 1
	c.sum += uint64(p[0])

	return l, nil
}

func (c AgeAverager) String() string {
	return fmt.Sprintf("%v", c.sum/c.count)
}
