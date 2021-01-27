package main

import (
  "fmt"
  "time"
  "os"
  "log"
  "bufio"
  "strings"
)

func getInput(input chan string) {
  for {
    in := bufio.NewReader(os.Stdin)
    result, err := in.ReadString('\n')
    if err != nil {
      log.Fatal(err)
    }
    input <- strings.TrimSuffix(result, "\n")
  }
}

func main() {
  answer := "74"
  fmt.Println("You have 5 seconds to calculate 19 * 4")
  input := make(chan string, 1)
  go getInput(input)

  for {
    select {
    case <- time.After(5 * time.Second):
      fmt.Println("Time's up!")
      return
    case i := <-input:
      if i == answer {
        fmt.Println("Correct.")
        return
      } else {
        fmt.Println("Incorrect.")
      }

    }
  }
}
