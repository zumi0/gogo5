package main

import "fmt"

func main() {
  i,j := 42, 2701
  p := &i
  *p = 21
  fmt.Println(i)
  q := &j
  *q = 32
  j = 21
  fmt.Println(j)
}
