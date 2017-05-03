package main

import (
	"fmt"
	"io"

	"github.com/another/c10-io.package/c10s11/ms"
)

func main() {
	fmt.Println("Demonstrates io.Copy")
	var n int
	var err error
	var data0 = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var mSrc ms.MemStore
	mSrc.Open(200)

	var mDest ms.MemStore
	mDest.Open(100)

	n, err = mSrc.Write(data0)
	fmt.Println("mSrc:", mSrc)

	io.Copy(&mDest, &mSrc)

	fmt.Println(n, err)
	fmt.Println("mDest:", mDest)

}
