package main

import "fmt"
import "github.com/another/c10-io.package/c10s10/ms"

func main() {
	fmt.Println("Demonstrates MemStore v2")
	var n int
	var err error
	var w0 = []byte{1, 2, 3, 4, 5, 6}
	var w1 = []byte{7, 8, 9, 10}

	var m ms.MemStore

	m.Open(7)

	n, err = m.Write(w0)
	fmt.Println(n, err)

	n, err = m.Write(w1)
	fmt.Println(n, err)

	fmt.Println(m)

	var r0 = make([]byte, 40)

	n, err = m.Read(r0)
	if n > 0 {
		fmt.Println(n, r0[:n])
	}

	if err != nil {
		fmt.Println(err)
	}

}
