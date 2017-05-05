package main

import (
	"fmt"

	"io"

	"github.com/another/c10-io.package/c10s14/ms"
)

func main() {
	fmt.Println("Example of io.TeeReader()")

	mSrc := ms.New(100)
	io.WriteString(mSrc, "It is a great day!")

	mDest := ms.New(200)
	r := io.TeeReader(mSrc, mDest)

	fmt.Println(mDest)

	buf := make([]byte, 100)
	r.Read(buf)

	fmt.Println(buf)
	fmt.Println(mDest)
}
