package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	// print in format YYYY-MM-DD hh:mm:ss
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05")) // full layout: "2006-01-02 15:04:05 -0700 MST"
}
