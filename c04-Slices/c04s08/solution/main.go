package main

import "fmt"

func main() {
	fmt.Println("Temperature Over 10 days")
	data := []int{14, 24, 37, 43, 49, 36, 33, 36, 40, 48}

	hhist(data)
	hist(data, 5)
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
func hist(d []int, sf int) {
	fmt.Println("\nHistogram")
	if 0 == len(d) {
		return
	}

	max := d[0]
	for i := 1; i < len(d); i++ {
		if max < d[i] {
			max = d[i]
		}
	}

	for i := max / sf; i > -1; i-- {
		fmt.Print("   ")
		for _, v := range d {
			if i <= v/sf {
				fmt.Print("*  ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

	// TODO - 1 : Print the value of each column
	fmt.Print("   ")
	for _, v := range d {
		fmt.Print(v, " ")
	}
	fmt.Println()
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
