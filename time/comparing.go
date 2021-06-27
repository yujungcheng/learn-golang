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

  // add
  old_tomorrow := tomorrow.Add(24 * time.Hour)
  new_today := time.Now()
  old_today := new_today.Add(-24 * time.Hour)
  new_tomorrow := new_today.Add(24 * time.Hour)

  fmt.Println("Old today: ", old_today)
  fmt.Println("Old tomorrow: ", old_tomorrow)
  fmt.Println("New today: ", new_today)
  fmt.Println("New tomorrow: ", new_tomorrow)
  
}
