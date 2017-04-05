package main

import (
	"fmt"
)

const C = 0x90909349298203829384293402989482904829048203948293429320492

type T chan int

func main() {

	var p *int
	var f func(int) string

	var m map[string]float32
	var s []complex64
	var ch chan T

	p = nil
	f = nil
	m = nil
	s = nil
	ch = nil

	var i float64 = C

	fmt.Println(p, f, m, s, ch, i)
}
