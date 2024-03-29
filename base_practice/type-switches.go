package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    //fmt.Printf("%s is %v bytes long\n", v, len(v))
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Println("Else type!")
  }
}

func main() {
  do(21)
  do("world")
  do(true)
}
