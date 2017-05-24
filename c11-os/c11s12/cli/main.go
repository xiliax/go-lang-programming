package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Standard Error : os.Stderr")

	os.Stdout.WriteString("This is going to the Standard Out (os.Stdout)\n")
	os.Stderr.WriteString("This is going to the Standard Error (os.Stderr)\n")
}
