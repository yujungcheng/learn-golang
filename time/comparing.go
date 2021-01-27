package main

import (
  "fmt"
  "time"
  "log"
)

func main() {
  t1 := "2021-01-02T15:04:05+07:00"
  t2 := "2021-01-03T15:04:05+07:00"

  today, err := time.Parse(time.RFC3339, t1)
  if err != nil {
    log.Fatal(err)
  }

  tomorrow, err := time.Parse(time.RFC3339, t2)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(today, "is after", tomorrow, "=>",today.After(tomorrow))
  fmt.Println(today, "is before", tomorrow, "=>", today.Before(tomorrow))
  fmt.Println(today, "is equal", tomorrow, "=>", today.Equal(tomorrow))
}
