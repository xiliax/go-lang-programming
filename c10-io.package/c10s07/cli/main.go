package main

import "fmt"
import "github.com/another/c10-io.package/c10s07/ms"

func main() {
	fmt.Println("Demo MemStore")

	var writeBuf = []byte{1, 2, 3, 4, 5}

	var m ms.MemStore

	m.Write(writeBuf)

	var readBuf = make([]byte, 20)
	m.Read(readBuf)
	fmt.Println(m)
	fmt.Println(readBuf)
}
