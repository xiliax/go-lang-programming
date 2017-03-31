package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Request struct {
	q string
	c chan<- Response
}

type Response struct {
	q        string
	a        string
	d        time.Duration
	serverId int
}

// TODO - create ClientResponse struct type (question, answer)

const N = 5
const MAX_SERVER_TIME = 100

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	qCh := client(Questions)
	ansCh := hub(qCh)

	for a := range ansCh {
		fmt.Println(a)
	}
}

func hub(qCh <-chan string) <-chan string {
	var wg sync.WaitGroup

	// TODO - make 'out' a chan ClientResponse
	out := make(chan string)

	go func() {
		for q := range qCh {
			ansCh := make(chan Response)
			req := Request{q: q, c: ansCh}

			for i := 0; i < N; i++ {
				go server(i, req)
			}

			wg.Add(1)
			go handler(out, q, ansCh, &wg)
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func handler(out chan<- string, theQuestion string, myAnsCh <-chan Response, wg *sync.WaitGroup) {
	timeout := time.After(MAX_SERVER_TIME / 2 * time.Nanosecond)

	select {
	case ans := <-myAnsCh:
		out <- ans.q + ":" + ans.a // TODO - create ClientResponse
	case <-timeout:
		out <- theQuestion + ":" + "NO SERVER RESPONSED" // TODO - create ClientResponse
	}
	wg.Done()
}

func server(id int, req Request) {
	d := time.Duration(rand.Intn(MAX_SERVER_TIME))
	time.Sleep(d * time.Nanosecond)
	a := Answers[req.q]
	resp := Response{q: req.q, a: a, d: d, serverId: id}
	req.c <- resp
}

func client(data []string) <-chan string {
	out := make(chan string)

	go func() {
		for _, q := range data {
			out <- q
		}
		close(out)
	}()

	return out
}
