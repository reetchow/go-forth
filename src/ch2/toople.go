package main

import "fmt"

func main() {
	for x := 0; x < 100; x++ {
		fmt.Printf("%d\t%d\n", x, fib(x))
	}
}

func fib(n int) int64 {
	// fib(n) = fib(n - 1) + fib(n - 2)
	var x, y int64 = 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
