package main

import (
    "fmt"
)

func main() {
    s := []int{}
    for i := 0; i <= 32; i++ {
    	s = appendInt(s, i)
        fmt.Printf("%v, len = %d, cap = %d\n", s, len(s), cap(s))
    }
}

func appendInt(s []int, n int) []int {
    appendLen := len(s) + 1
    if appendLen <= cap(s) {
        s[appendLen - 1] = n
        return s[:appendLen]
    } else {
	var newCapacity int
	if cap(s) == 0 {
	    newCapacity = 1
	} else {
	    newCapacity = 2 * cap(s)
	}
	z := make([]int, appendLen, newCapacity)
        copy(z, s)
        z[len(s)] = n
        return z
    }
}
