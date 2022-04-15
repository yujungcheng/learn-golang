package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	s := "2021-01-02T15:04:05+07:00"

	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Original time        :", t)

	// Add/substrate time
	addT := t.Add(2 * time.Second)
	fmt.Println("Add 2 seconds        :", addT)
	subT := t.Add(-2 * time.Second)
	fmt.Println("Del 2 seconds        :", subT)

	// AddDate
	fmt.Println("Add 1 day            :", t.AddDate(0, 0, 1))
	fmt.Println("Add 1 month          :", t.AddDate(0, 1, 0))
	fmt.Println("Add 1 year and 1 day :", t.AddDate(1, 0, 1))
	fmt.Println("Del 1 year           :", t.AddDate(-1, 0, 0))

	// Duration
	n := "2021-01-20T22:08:10+07:00"
	p, err := time.Parse(time.RFC3339, n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nNew time             :", t)
	fmt.Println("Diff with orig time  :", p.Sub(t))
	fmt.Println("Diff with new time   :", t.Sub(p))
}
