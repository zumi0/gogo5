package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) float64 {
  z := float64(1)
  for i := 0; i < 10; i++ {
    z = (z*z + x)/(2*z)
  }
  return z
}

func main() {
  error := math.Abs(math.Sqrt(2)-sqrt(float64(2)))
  fmt.Println("Dedault:", math.Sqrt(2))
  fmt.Println("My function:", sqrt(float64(2)))
  fmt.Println("Error:", error)
}
