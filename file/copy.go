package main

import (
  "fmt"
  "log"
  "os"
  "io"
)

func main() {
  fromFile := "./new_file.txt"
  toFile := "old_file.txt"
  from, err := os.Open(fromFile)
  if err != nil {
    log.Fatal(err)
  }
  defer from.Close()  // close file after all other execution finished

  to, err := os.OpenFile(toFile, os.O_RDWR|os.O_CREATE, 0666)
  if err != nil {
    log.Fatal(err)
  }
  defer to.Close()  // close file after all other execution finished

  fmt.Printf("Copy files from %v to %v \n", fromFile, toFile)
  _, err = io.Copy(to, from)
  if err != nil {
    log.Fatal(err)
  }
}
