package main

func main() {
	aFunc()
	go foo()
	go hoo()
}

func foo() {
	aFunc()
}

func hoo() {
	aFunc()
}

func aFunc() {

}
