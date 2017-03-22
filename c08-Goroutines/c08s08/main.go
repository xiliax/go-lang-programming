package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	chatter("Bob", &wg)
	chatter("Joe", &wg)
	chatter("Anne", &wg)
	chatter("Mary", &wg)
	chatter("Jane", &wg)
	chatter("Mike", &wg)
	chatter("Peter", &wg)

	wg.Wait()

	fmt.Println("All chatter have finished")
}

func chatter(s string, wg *sync.WaitGroup) {
	talkTime := rand.Intn(3000)
	timeUp := time.After(time.Duration(talkTime) * time.Nanosecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-timeUp:
				// clean up
				fmt.Println(s, " is quitting after", time.Duration(talkTime))
				return
			default:
				fmt.Println(s)
			}
		}
	}()

}
