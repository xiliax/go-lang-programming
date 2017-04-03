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
type ClientResponse struct {
	q string
	a string
}
type ClientRequest struct {
	q string
	c chan<- ClientResponse
}

const N = 5
const MAX_SERVER_TIME = 100

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	qCh := client(Questions)
	ansCh := hub(qCh)

	for a := range ansCh {
		fmt.Printf("Questions: %v\n\t--> %v\n", a.q, a.a)
	}
}

func hub(qCh <-chan string) <-chan ClientResponse {
	var wg sync.WaitGroup

	// TODO - make 'out' a chan ClientResponse
	out := make(chan ClientResponse)

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

func handler(out chan<- ClientResponse, theQuestion string, myAnsCh <-chan Response, wg *sync.WaitGroup) {
	timeout := time.After(MAX_SERVER_TIME / 2 * time.Nanosecond)

	cr := ClientResponse{q: theQuestion}

	select {
	case ans := <-myAnsCh:
		cr.a = ans.a
		out <- cr // TODO - create ClientResponse
	case <-timeout:
		cr.a = "NO SERVER RESPONSED"
		out <- cr // TODO - create ClientResponse
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

func client(data []string) <-chan ClientRequest {
	out := make(chan ClientRequest)
	myResponse := make(chan ClientResponse)

	go func() {
		for _, q := range data {
			out <- ClientRequest{q: q, c: myResponse}
		}
		close(out)
	}()

	return out
}
