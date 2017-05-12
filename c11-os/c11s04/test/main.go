package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Using x.Stat()")
	fn := "awesome"
	fi, err := os.Stat(fn)

	if nil != err {
		log.Fatalf("%v\n", err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("is Directory")
	case mode.IsRegular():
		fmt.Println("is regular file")
	default:
		fmt.Println("is something else")
	}
}
