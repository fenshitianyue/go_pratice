package main

import "fmt"

func sum(s []int, c chan int) {
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum // 将 和 送入channel c
}

func main() {
  s := []int{7, 2, 1, 9, -3, 5}
  c := make(chan int)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  x, y := <-c, <-c  // 从 channel c 中接收
  fmt.Println(x, y, x+y)
}
