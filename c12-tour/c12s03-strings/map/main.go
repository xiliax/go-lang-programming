package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Using strings.Map")
	s := "abc-efg"
	result := strings.Map(mapFunc, s)
	fmt.Printf("in: %s => out: %s\n", s, result)
}

func mapFunc(in rune) rune {
	switch in {
	case 'a':
		return 'e'
	case 'b':
		return 'f'
	case 'c':
		return 'g'
	case 'e':
		return 'a'
	case 'f':
		return 'b'
	case 'g':
		return 'c'
	}

	return in
}
