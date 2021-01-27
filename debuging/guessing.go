package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)


func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Guess the name of my pet to win a prize: ")
  text, _ := reader.ReadString('\n')
  text = strings.Replace(text, "\n", "", -1)

  fmt.Println("[DEBUG] text is:", text)

  if text == "Laifu" {
    fmt.Println("You won!")
  } else {
    fmt.Println("You didn't win.")
  }  
}
