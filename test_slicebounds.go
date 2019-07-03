package main

import "fmt"

func main() {
  s := []int{1,2,3,4,5}
  t := s[2:4]
  u := s[:3]

  fmt.Println(s,t,u)
}
