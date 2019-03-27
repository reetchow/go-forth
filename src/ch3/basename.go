package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if args := os.Args[1:]; len(args) > 0 {
		s := args[0]
		fmt.Printf("%s -> %s\n", s, basename1(s))
		fmt.Printf("%s -> %s\n", s, basename2(s))
	}
}

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i + 1:]
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slashIdx := strings.LastIndex(s, "/")
	s = s[slashIdx + 1:]
	if dotIdx := strings.LastIndex(s, "."); dotIdx >= 0 {
		s = s[:dotIdx]
	}
	return s
}
