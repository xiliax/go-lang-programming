package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Iteration using 'for' Loop and range()")
	var m = map[int]int{0: 21, 1: 53, 3: 4, 5: 126, 6: 8}

	fmt.Println(m)
	v, ok := m[2]
	if ok {
		fmt.Println("m[2]=", v)
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("m[%v]=%v\n", k, m[k])
	}
}
