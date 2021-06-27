package main

import (
  "fmt"
  "io/ioutil"
  "errors"
)

func Half(numberToHalf int) (int, error) {
  if numberToHalf % 2 != 0 {
    return -1, fmt.Errorf("Cannot half %v", numberToHalf)
  }
  return numberToHalf / 2, nil
}

func main() {
  file, err := ioutil.ReadFile("foo.txt")
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("%s", file)
  }

  err = errors.New("something went wrong")
  if err != nil {
    fmt.Println(err)
  }

  // formatting errors
  name, role := "Richard Jupp", "Drummer"
  err = fmt.Errorf("The %v %v quit", role, name)
  if err != nil {
    fmt.Println(err)
    }

  n, err := Half(19)
  if err != nil {
    fmt.Println(err)
    //return
  } else {
    fmt.Println(n)
  }

  // using panic to halt execution
  fmt.Println("This is executed")
  panic("I can do no more. Goodbye.")  // execution is halted after panic is called
  fmt.Println("This is not executed.")  // never executed

}
