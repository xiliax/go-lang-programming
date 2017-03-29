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
	wordsCh1 := lineSplitter(lineCh)

	wg.Add(1)
	go wordCounter(&result, wordsCh0, wordsCh1, &wg)

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

func wordCounter(result *map[string]int, wordsCh0, wordsCh1 <-chan []string, wg *sync.WaitGroup) {
	inCh := make(chan []string)

	// we need to know when there are no more words from 'wordsCh0' and 'wordsCh1', so we create our own wait WaitGroup
	var myWg sync.WaitGroup

	// the next two goroutines does fan-in from 'wordsCh0' and 'wordsCh1' to 'inCh'
	myWg.Add(1)
	go func() {
		for words := range wordsCh0 {
			inCh <- words
		}
		myWg.Done()
	}()

	myWg.Add(1)
	go func() {
		for words := range wordsCh1 {
			inCh <- words
		}
		myWg.Done()
	}()

	// this goroutine does the word count by gettings words from 'inCh'
	go func() {
		for words := range inCh {
			for _, w := range words {
				(*result)[w] = (*result)[w] + 1
			}
		}
	}()

	myWg.Wait() // wait for the to first goroutines launched in this function to complete

	// at this point, both fan-in goroutines are finished reading words, so we can close 'inCh',
	// this will signal to our word counter goroutine to exit
	close(inCh)

	wg.Done()
}
