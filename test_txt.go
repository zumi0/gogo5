package main

import (
  "fmt"
  "os"
)

func main() {
  f, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer f.close()


}
