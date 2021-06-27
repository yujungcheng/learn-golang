package main

import (
  "fmt"
  "strings"
)

func main() {

  string1 := "aaa"
  string2 := "bbb"
  string3 := "aaaccc"
  string4 := "aaa"

  fmt.Println("string1 == string2: ", string1 == string2)
  fmt.Println("string3 != string4: ", string3 != string4)


  fmt.Println("GFG > Geeks: ", "GFG" > "Geeks")
  fmt.Println("GFG < Geeks: ", "GFG" < "Geeks")
  fmt.Println("GFG >= For: ", "GFG" >= "For")
  fmt.Println("GFG <= For: ", "GFG" <= "Geeks")

  fmt.Println("Compare: ", strings.Compare("gfg", "fgf"))
  fmt.Println("Compare: ", strings.Compare("gfg", "gfg"))
  fmt.Println("Compare: ", strings.Compare("gfg", "gfG"))
}
