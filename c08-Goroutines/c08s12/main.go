package main

import "fmt"
import "strings"
import "sync"

func main() {
	var result map[string]int
	result = make(map[string]int)
	var wg sync.WaitGroup

	lineCh := reader(Data)

	wordsCh0 := lineSplitter(lineCh)

	wg.Add(1)
	go wordCounter(&result, wordsCh0, &wg)

	wg.Wait()
	for k, v := range result {
		fmt.Println(v, "->", k)
	}
}

func reader(d []string) <-chan string {
	out := make(chan string)

	go func() {
		for _, line := range d {
			out <- line
		}
		close(out)
	}()

	return out
}

func lineSplitter(lineCh <-chan string) <-chan []string {
	out := make(chan []string)

	go func() {
		for l := range lineCh {
			out <- strings.Split(l, " ")
		}
		close(out)
	}()

	return out
}

func wordCounter(result *map[string]int, wordsCh <-chan []string, wg *sync.WaitGroup) {
	for words := range wordsCh {
		for _, w := range words {
			(*result)[w] = (*result)[w] + 1
		}
	}

	wg.Done()
}
