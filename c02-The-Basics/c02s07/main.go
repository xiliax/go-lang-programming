package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("Flow Controll and the 'if' Statement")
	now := time.Now()
	rand.Seed(now.Unix())
	age := rand.Intn(150)

	if age < 18 {
		fmt.Printf("You are %v years old, you are too young to sign up.\n", age)
	} else if age < 64 {
		fmt.Printf("Hey, awesome to you sign up\n")
	} else if age < 100 {
		fmt.Printf("Thank you seniors\n")
	} else {
		fmt.Println("You are awsome.")
	}
}
