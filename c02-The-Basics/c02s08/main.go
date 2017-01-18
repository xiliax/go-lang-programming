package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("Flow Controll and the 'switch' Statement")
	now := time.Now()
	rand.Seed(now.Unix())
	age := rand.Intn(150)

	switch {
	case age < 18:
		fmt.Println("Too young to sign up")
	case age < 65:
		fmt.Println("You are 18 to 64 years old")
	case age < 100:
		fmt.Println("You are not 100 yet")
	default:
		fmt.Println("You are over 100 years old")
	}
}
