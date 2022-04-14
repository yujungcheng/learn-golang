package main

import (
  "fmt"
)

func main () {
  b := true
  if b {
    fmt.Println("b is true!")
  } else  {
    fmt.Println("b is false!")
  }

  i := 4
  if i == 3 {
    fmt.Println("i is 3")
  } else if i == 2 {
    fmt.Println("i is 2")
  } else if i < 2 {
    fmt.Println("i less than 2")
  } else if i >= 4 {
    fmt.Println("i greater or equal 4")
  }

  j := 2
  switch j {
  case 2:
    fmt.Println("Two")
  case 3:
    fmt.Println("Three")
  default:
    fmt.Println("Unknown!")
  }

  k := 0
  for k < 10 {
    k++
    fmt.Println("k is", k)
  }

  for p := 0; p < 10; p++ {
    fmt.Println("p is", p)
  }

  numbers := []int{1, 2, 3, 4}
  for i, n := range numbers {
    fmt.Println("index of the loop is", i)
    fmt.Println("value from the array is", n)
  }

  // string iteration 
  for i, ch := range "a str" {
    fmt.Printf("%#U starts at byte position %d\n", ch, i)
  }
  const s = "b str"
  for i := 0; i < len(s); i++ {
    fmt.Printf("%x", s[i])
  }
  fmt.Println("")

  // channel iteration
  ch := make(chan int)
  go func() {
    ch <- 10
    ch <- 20
    ch <- 30
  }()
  for n := range ch {
    fmt.Println(n)
  }

  defer fmt.Println("I am run after the function completes")
  fmt.Println("Hello world")

  defer fmt.Println("I am the first defer")
  defer fmt.Println("I am the second defer")
  defer fmt.Println("I am the third defer")
  fmt.Println("Hello world2")
}
