package main

import (
  "io/ioutil"
  "log"
)

func main() {
  b := make([]byte, 0) // empty slice of byte
  err := ioutil.WriteFile("new_file.txt", b, 0644)
  if err != nil {
    log.Fatal(err)
  }
}
