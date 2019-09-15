package main

import (
  "fmt"
)

func main() {
  a := 10
  fmt.Printf("十进制: %d\n", a)
  fmt.Printf("二进制: %b\n", a)
  fmt.Printf("八进制: %#o\n", a)
  fmt.Printf("十六进制: %#x\n", a)
}
