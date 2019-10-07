package main

import "fmt"

type Person struct {
  Name string
  Age int
}

// func (p Person) String() string {
//   return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
// }

func main() {
  a := Person{"zanda", 30}
  b := Person{"aideny", 9000}
  fmt.Println(a, b)
}

