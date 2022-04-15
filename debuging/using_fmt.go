package main

import (
	"fmt"
)

type Animal struct {
	Name  string
	Color string
}

func main() {
	a := Animal{
		Name:  "Cat",
		Color: "Black",
	}

	fmt.Printf("Value: %v\n", a)
	fmt.Printf("Field Name and Value: %+v\n", a)
}
