package main

import (
  "fmt"
)

func main() {
  board := [][]string{
    []string{"_","_","_"},
    []string{"_","_","_"},
    []string{"_","_","_"},
  }

  board[0][0] = "X"
  fmt.Println(board[0])
}
