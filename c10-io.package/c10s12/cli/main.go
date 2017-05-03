package main

import (
	"fmt"
	"io"

	"github.com/another/c10-io.package/c10s12/ms"
	"github.com/another/c10-io.package/c10s12/p"
)

func main() {
	fmt.Println("The Problem with io.Copy and Person")

	bob := p.Person{"Bobby Jones", 35, "100-01-1101"}
	var jane p.Person

	var m ms.MemStore
	m.Open(200)

	n, err := io.Copy(&m, bob)
	n, err = io.Copy(&jane, &m)
	fmt.Println("jane:", jane)

	fmt.Println(n, err)

}
