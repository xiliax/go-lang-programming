package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Line struct {
	length int
	text   string
}

func main() {
	if 1 < len(os.Args) {
		f, _ := os.Open(os.Args[1])
		defer f.Close()

		var t string
		list := []Line{}
		scanner := bufio.NewScanner(f)
		maxLengthWidth := 0
		maxLength := 0
		for scanner.Scan() {
			t = scanner.Text()
			maxLength = len(t)
			list = append(list, Line{maxLength, t})
			maxLengthWidth = max(maxLengthWidth, len(strconv.Itoa(maxLength)))
		}
		for _, l := range list {
			fmt.Printf("%*v | %v\n", maxLengthWidth, l.length, l.text)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Please enter a filename")
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
