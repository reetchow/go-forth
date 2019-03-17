package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	if len(os.Args) > 2 {
		for i := 2; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = os.Args[1]
		}
		fmt.Println(s)
	}
	fmt.Println(time.Since(start).Nanoseconds())
}
