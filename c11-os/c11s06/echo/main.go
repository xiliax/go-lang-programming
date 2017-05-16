package main

import (
	"fmt"
	"os"
)

var printNewLine = "\n"

/* The echo utility writes any specified operands, separated by single blank
   (` ') characters and followed by a newline (`\n') character, to the stan-
   dard output.

   -n    Do not print the trailing newline character.
*/
func main() {
	if 1 == len(os.Args) {
		fmt.Println("")
		return
	}
	offset := 1 // skip program name
	if "-n" == os.Args[1] {
		printNewLine = ""
		offset++ // skip -n option
	}

	outputSep := ""
	for _, a := range os.Args[offset:] {
		fmt.Printf("%v%v", outputSep, a)
		outputSep = " "
	}

	fmt.Print(printNewLine)
}
