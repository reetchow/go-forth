package main

import (
  "fmt"
)

func main() {
  var s []int = []int{0}
  for i := 1; i <= 32; i++ {
    s = append(s, i)
    fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
  }
  s = []int{0}
  for i := 1; i <= 32; i++ {
    s = appendInt(s, i)
    fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
  }
}

func appendInt(s []int, n int) []int {
  zlen := len(s) + 1
  if zlen <= cap(s) {
    z := s[:zlen]
    z[zlen - 1] = n
    return z
  } else {
    zcap := 2 * cap(s)
    z := make([]int, zlen, zcap)
    copy(z, s)
    z[zlen - 1] = n
    return z
  }
}
