package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("Flow Controll and the 'for' Statement")
	now := time.Now()
	rand.Seed(now.Unix())

	var age int
	count := 1
	for ; ; count++ {
		age = rand.Intn(150)

		if age < 25 {
			fmt.Println("Age is less than 25, so starting over -- ", count)
			continue
		}

		if age < 100 {
			fmt.Println("Age", age, "you are not 100 yet -- ", count)
		} else {
			fmt.Println("Age", age, "you are over 100 years old -- ", count)
			fmt.Println("Exiting the loop early")
			break
		}

	}
}
