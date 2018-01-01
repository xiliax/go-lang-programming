package main

import "fmt"

func main() {
	var buffer *byte
	var pBuffer = &buffer
	var message = "Some very long message"
	var pMessage = &message // where is the very long message in memory?
	var isLoggingEnabled = true
	var pLoggingEnabled = &isLoggingEnabled

	var count = 1737
	var addressOfCount = &count
	var pointerToPointerToCount = &addressOfCount

	fmt.Println("--- Address of, &, operator examples:")
	fmt.Printf("buffer : %#v\n", buffer)
	fmt.Printf("pBuffer : %#v\n", pBuffer)
	fmt.Printf("pMessage: %#v\n", pMessage)
	fmt.Printf("ptrLoggingEnabled: %#v\n", pLoggingEnabled)
	fmt.Printf("count : %#v\n", count)
	fmt.Printf("addressOfCount : %#v\n", addressOfCount)
	fmt.Printf("pointerToPointerToCount: %#v\n", pointerToPointerToCount)
}
