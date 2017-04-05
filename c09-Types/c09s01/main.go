package main

import "fmt"
import "time"

type StudentId int

func main() {
	fmt.Println("Defining New Types")
	student0 := StudentId(25)

	printStudentId(25)
	printStudentId(student0)
	id := 25
	fmt.Printf("id: %T -> %v\n", id, id)

	// printStudentId(id)
	time.Sleep(25 * time.Nanosecond)
	// time.Sleep(time.Duration(id) * time.Nanosecond)
}

func printStudentId(id StudentId) {
	fmt.Printf("Student0: %T -> %v\n", id, id)
}
