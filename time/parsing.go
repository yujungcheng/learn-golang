package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	timeString := "2021-01-02T15:04:05+07:00"
	fmt.Println("Original time string: ", timeString, "\n")

	t, err := time.Parse(time.RFC3339, timeString) // parse into Time struct.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	fmt.Printf("hour is %v\n", t.Hour())
	fmt.Printf("minute is %v\n", t.Minute())
	fmt.Printf("second is %v\n", t.Second())
	fmt.Printf("day is %v\n", t.Day())
	fmt.Printf("month is %v (%v)\n", t.Month(), int(t.Month()))
	fmt.Printf("UNIX time is %v\n", t.Unix())
	fmt.Printf("week day is %v (%v)\n", t.Weekday(), int(t.Weekday()))
}
