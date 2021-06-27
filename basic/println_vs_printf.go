package main

import "fmt"

func main() {

  // Println
  x := "3"
  fmt.Println("1 +", x, "= 4")

  a, b := 4, 5
  fmt.Println(a, "x", b, "=", a*b)


  // Printf
  c := "5"
  d := 10
  fmt.Printf("2 + 3 = %s", c)
  fmt.Printf(" and %s + 3 = %d", c, d)
  fmt.Printf(
    "\n%s x %d = %d\n", c, 2, d,
  )
}
