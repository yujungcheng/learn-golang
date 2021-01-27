package main

import (
  "fmt"
  "strconv"
  "bytes"
  "strings"
)

func main() {
  s := "I am an interpreted string literal"
  fmt.Println(s)

  rune_s := "After a backslash, certain single character escapes represent special values\nn is a line feed or new line \n\t t is a tab"
  fmt.Println(rune_s)
  raw_s := `After a backslash, certain single character escapes represent special values`
  fmt.Println(raw_s)

  concatenated_s := "Oh sweet ignition" + " be my fuse"
  fmt.Println(concatenated_s)

  a := "Can you hear me?"
  a += "\nHear me screamin?"
  fmt.Println(a)

  var i int = 1
  var egg_s string = " egg"
  int_to_string := strconv.Itoa(i)
  var breakfast string = int_to_string + egg_s
  fmt.Println(breakfast)

  // buffer
  var buffer bytes.Buffer
  for i := 0; i < 100; i++ {
    buffer.WriteString("z")
  }
  fmt.Println(buffer.String())

  fmt.Println(strings.ToLower("MESSAGE"))

  fmt.Println(strings.Index("surface", "face"))  // found, return position index
  fmt.Println(strings.Index("moon", "aer"))  // not found, return -1
  fmt.Println(strings.TrimSpace(" I don't need all this space "))  // Trimming space

}
