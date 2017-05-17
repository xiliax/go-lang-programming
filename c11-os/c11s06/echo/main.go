package main

import (
	"fmt"
	"os"
)

/*
echo -- write arguments to the standard output

The echo utility writes any specified operands, separated by single blank (` ')
characters and followed by a newline (`\n') character, to the standard output.

     The following option is available:

     -n    Do not print the trailing newline character.
*/

var lastChar = "\n" // defaults to a \n, unless -n is specified
func main() {
	// skip over program name and possible -n option
	offset := 1 // skips program name
	if 2 <= len(os.Args) && "-n" == os.Args[offset] {
		lastChar = ""
		// if we are in here, we need to skip -n option
		offset++
	}

	// write each remaining args to output with ' ' between them
	for _, a := range os.Args[offset:] {
		fmt.Print(a, " ")
	}

	fmt.Print(lastChar)
}
