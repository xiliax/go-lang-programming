package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

func main() {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			fmt.Println("My id is ", id)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
