package main

import (
	"fmt"
	"io"

	"github.com/another/c10-io.package/c10s13/ms"
)

func main() {
	fmt.Println("io.MultiReader and MemStore v3")

	m0 := ms.New(0)
	io.WriteString(m0, "Hello, World!")
	fmt.Println(m0)

	m1 := ms.New(-10)
	io.WriteString(m1, "It is a wonderful day!")
	fmt.Println(m1)

	m2 := ms.New(200)
	io.WriteString(m2, "I love Go Language!")
	fmt.Println(m2)

	m3 := ms.New(200)

	mr := io.MultiReader(m0, m1, m2)
	io.Copy(m3, mr)

	fmt.Println(m3)
}
