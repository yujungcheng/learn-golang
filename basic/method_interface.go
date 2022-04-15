package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

// method sets
type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

type Triangle struct {
	base   float64
	height float64
}

// passing pointer reference to method
func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

// passing a value reference to method
func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

// passing pointer reference to method
func (t *Triangle) changeBase2(f float64) {
	t.base = f
	return
}

// interfaces
type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name   string
	Broken bool
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Name   string
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	m := Movie{
		Name:   "spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())

	t := Triangle{base: 3, height: 1}
	fmt.Println(t.area())

	p := Triangle{base: 3, height: 1}
	p.changeBase(4)
	fmt.Println(p.base)

	q := Triangle{base: 3, height: 1}
	q.changeBase2(4)
	fmt.Println(q.base)

	// interfaces
	t850 := T850{
		Name:   "The Terminator",
		Broken: false,
	}

	r2d2 := R2D2{
		Name:   "R2D2",
		Broken: true,
	}

	err := Boot(&r2d2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}

	err = Boot(&t850)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}
