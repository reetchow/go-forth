package main

import (
	"fmt"
	"os"
)

func main() {
	if args := os.Args[1:]; len(args) > 0 {
		s := args[0]
		fmt.Printf("%s -> %s", s, commas(s))
	}
}

func commas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commas(s[:n - 3]) + "," + s[n - 3:]
}
