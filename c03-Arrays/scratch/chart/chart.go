package main

import "fmt"

func main() {
	data := []int{14, 24, 37, 43, 49, 36, 33, 36, 40, 38}
	hhist(data)
	hist(data[:])
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
	chart := make([][]int, len(d))

	// initialize each horizontal line of the chart
	for i := 0; i < len(d); i++ {
		chart[i] = make([]int, d[i])
	}

	// display chart
	for i := 0; i < len(d); i++ {
		fmt.Print(len(chart[i]), " ")
		for j := 0; j < len(chart[i]); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
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
func hist(d []int) {
	if len(d) == 0 {
		return
	}

	max := d[0]
	for i := 1; i < len(d); i++ {
		if max < d[i] {
			max = d[i]
		}
	}

	for i := max / 3; i > -1; i-- {
		fmt.Print("  ")
		for j := 0; j < len(d); j++ {
			if i <= (d[j] / 3) {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

}
