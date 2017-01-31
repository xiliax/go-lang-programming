package main

import (
	"fmt"
)

func main() {
	fmt.Println("Introduciton to Maps")
	var m map[int]int
	var m1 = map[int]int{0: 21, 1: 53, 2: 73, 3: 4, 5: 126, 6: 8}
	var m2 = map[int]float64{0: 21.5, 1: 0.053, 2: -1.73, 3: 4, 5: 126.93, 6: 8}
	var m3 = map[int]bool{1: true, 5: true, 4: false, 10: false, 101: true, 99: false}

	m4 := map[int]string{0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday",
		5: "Friday", 6: "Saturday"}

	m5 := map[string]int{"Sunday": 0, "Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4,
		"Friday": 5, "Saturday": 6}

	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println(m2[1])
	fmt.Println(m3)
	fmt.Println(m4)
	day := "Thursday"
	fmt.Println(day, "=>", m5[day])

}
