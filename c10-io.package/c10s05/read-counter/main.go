package main

import "fmt"

type ReadCounter struct {
	count uint64
}

func main() {
	fmt.Println("Read Counter")
	var rc ReadCounter
	var s0 []byte
	var s1 = make([]byte, 20)
	var s2 = make([]byte, 56)

	rc.Read(s0)
	rc.Read(s1)
	rc.Read(s0)
	rc.Read(s1)
	rc.Read(s2)

	fmt.Printf("%v\n", rc)
}

func (r *ReadCounter) Read(p []byte) (n int, err error) {
	r.count++
	return len(p), nil
}

func (r ReadCounter) String() string {
	return fmt.Sprintf("%v", r.count)
}
