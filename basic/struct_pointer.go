package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Name   string
	Rating float32
}

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	// Create a struct variable, method 1 (prefer)
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}
	fmt.Println(m.Name, m.Rating)

	// Create a struct variable, method 2
	var n Movie
	fmt.Printf("%+v\n", n)
	n.Name = "Metropolis"
	n.Rating = 0.99
	fmt.Printf("%+v\n", n)

	// Create a struct variable, method 3
	o := new(Movie)
	o.Name = "Metropolis"
	o.Rating = 0.99
	fmt.Printf("%+v\n", o)

	c := Movie{Name: "Citizen Kane2", Rating: 20}
	fmt.Printf("%+v\n", c)

	// nesting structs
	e := Superhero{
		Name: "Batman",
		Age:  32,
		Address: Address{
			Number: 1007,
			Street: "Mountain Drive",
			City:   "Gotham",
		},
	}
	fmt.Printf("%+v\n", e)
	fmt.Println(e.Address.Street)

	fmt.Printf("%+v\n", NewAlarm("07:00"))

	// comparing structs
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	if a == b {
		fmt.Println("a and b are the same")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)

	fmt.Printf("%+v\n", reflect.TypeOf(a))
	fmt.Printf("%+v\n", reflect.TypeOf(b))

	// copy a struct using a value reference
	d := a
	d.Ice = false
	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &d)

	// copy a struct using a pointer reference
	f := &a // assign by pointer
	f.Ice = false
	fmt.Printf("%+v\n", *f)
	fmt.Printf("%+v\n", f) // print '&'' which means it store address which point to another variable
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", f)
	fmt.Printf("%p\n", &a)
}
