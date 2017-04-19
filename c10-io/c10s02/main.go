package main

import "fmt"

type Counter int

func main() {
	fmt.Println("Bytes Counter")

	var c Counter
	var d0 = []byte("Hello, world!")
	var d1 []byte
	var d2 []byte

	c.Write(d0)
	c.Write([]byte{0, 9, 10})
	c.Write(d1)
	c.Write(d2)
	c.Write([]byte{0, 5, 9, 10})

	fmt.Printf("Bytes Counter: %v\n", c)
}

func (c *Counter) Write(p []byte) (n int, err error) {
	l := len(p)
	*c = Counter(int(*c) + l)
	return l, nil
}
