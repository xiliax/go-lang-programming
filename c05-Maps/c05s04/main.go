package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps Review - Histogram")
	data := map[int]int{2: 3, 3: 6, 4: 9, 5: 13, 6: 15, 7: 18,
		8: 16, 9: 14, 10: 10, 12: 2}

	hist(data, 1)
}

/*
 Histogram
10    *
9     *
8     *  *
7     *  *     *
6     *  *     *              *
5     *  *  *  *        *     *
4     *  *  *  *        *     *
3     *  *  *  *  *     *     *
2     *  *  *  *  *     *     *  *
1     *  *  *  *  *     *  *  *  *
0     *  *  *  *  *  *  *  *  *  *
      1  2  3  4  5  6  7  8  9  10
*/
func hist(d map[int]int, sf int) {
	fmt.Println("\nHistogram")
	if 0 == len(d) {
		return
	}

	max := 0
	keys := []int{}
	for k, v := range d {
		keys = append(keys, k)
		if max < v {
			max = v
		}
	}

	sort.Ints(keys)

	for i := max / sf; i > -1; i-- {
		fmt.Print("  ")
		for _, k := range keys {
			if i <= d[k]/sf {
				fmt.Print("*  ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------------------")
}

/*
 Histogram (horizontal)
 ************************
 *****************
 ********************
 *
 ************
 ******
 *****************
 ***********************
*/
func hhist(d []int) {
	fmt.Println("\nHistogram (horizontal)")
	for _, v := range d {
		fmt.Print(v, " ")
		for j := 0; j < v; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------------------")
}
