package main

import "fmt"

func main() {
	var buffer *byte
	var pBuffer = &buffer
	var message = "Some very long message"
	var pMessage = &message
	var isLoggingEnabled = true
	var pLoggingEnabled = &isLoggingEnabled

	var count = 1737
	var addressOfCount = &count
	var pointerToPointerToCount = &addressOfCount

	fmt.Println("--- Dereference, *, operator examples:")
	// fmt.Printf("*buffer : %#v\n", *buffer) 	// <<<----- this will case the program to crash
	fmt.Printf("*pBuffer : %#v\n", *pBuffer)  // address of buffer : nil
	fmt.Printf("*pMessage: %#v\n", *pMessage) // -> Some very long message"
	fmt.Printf("*pLoggingEnabled: %#v\n", *pLoggingEnabled)
	//fmt.Printf("count : %#v\n", *count)		// not a pointer, can't dereference
	fmt.Printf("*addressOfCount : %#v\n", *addressOfCount)                    // -> count : 1737
	fmt.Printf("**pointerToPointerToCount: %#v\n", **pointerToPointerToCount) // -> count : 1737
}
