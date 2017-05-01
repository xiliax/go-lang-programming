package main

import "fmt"
import "github.com/another/c10-io.package/c10s09/ms"

func main() {
	fmt.Println("io Errors")

	var w0 = []byte{1, 2, 3, 4, 5, 6}
	var w1 = []byte{7, 8, 9, 10}

	var m ms.MemStore

	m.Write(w0)
	m.Write(w1)

	fmt.Println(m)

	var r0 = make([]byte, 40)
	var n int
	var err error

	n, err = m.Read(r0)
	if n > 0 {
		fmt.Println(n, r0[:n])
	}

	if err != nil {
		fmt.Println("End of stream! No more data to read")
	}

}
