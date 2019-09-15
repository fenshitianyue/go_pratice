package main

import (
  "fmt"
)

func main() {
  s0 := "hello"
  s0 += " "
  s1 := "world"
  s2 := s0 + s1
  fmt.Printf(s0 + "\n")
  fmt.Printf(s1 + "\n")
  fmt.Printf(s2 + "\n")
}

