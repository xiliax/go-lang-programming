package main

import "fmt"

type ServerResponse struct {
	q string
	a string
}

type ClientResponse struct {
	q string
	a string
}

type StudenId int

type MyFunc func(i int, s string)

func main() {
	fmt.Println("Chapter 9: Types (preview)")

	var sr = ServerResponse{"another questions", "another answer"}

	var cr = ClientResponse{"question", "answer"}

	printQuestion(sr)
	// printQuestion(cr)
	fmt.Println(cr)

	var x StudenId = 5
	printStudentId(x)

	var myFun MyFunc = func(i int, s string) {
		// ...
	}
	go myFun(5, "name")

	myFun = func(x int, z string) {

	}
	go myFun(10, "other call")

	goo(myFun)

}
func printStudentId(id StudenId) {
	fmt.Println(id)
}
func printQuestion(req ServerResponse) {
	fmt.Println("Questions: ", req.q)
}

func goo(f MyFunc) {
	f(42, "life of the universe")
}
