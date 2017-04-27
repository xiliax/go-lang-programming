package main

import "fmt"
import "github.com/another/c10-io.package/c10s08/ms"

func main() {
	fmt.Println("Demo MemStore")

	var w0 = []byte{1, 2, 3, 4, 5, 6}
	var w1 = []byte{7, 8, 9, 10}

	var m ms.MemStore

	m.Write(w0)
	m.Write(w1)

	fmt.Println(m)

	var r0 = make([]byte, 4)
	var n int
	var err error
	for err == nil {
		n, err = m.Read(r0)
		fmt.Println(n, err, r0)
	}

	m.Reset()
	n, err = m.Read(r0)
	fmt.Println(n, err, r0)

	m.Close()

	n, err = m.Read(r0)
	if err != nil {
		fmt.Println(n, err)
		return
	}

	fmt.Println(n, err, r0)
}
