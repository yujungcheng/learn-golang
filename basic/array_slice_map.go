package main

import (
	"fmt"
)

func main() {
	// array
	var a [3]int //int array with length 3
	a[0] = 12    // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(b)

	// colon : is used when you perform the short declaration and assignment for the first time as you are doing in your first statement i.e. myArray  :=[...]int{12,14,26}

	// slices
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses_de_Bourgogne"

	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])

	cheeses = append(cheeses, "Camembert")
	fmt.Println(cheeses[2])

	cheeses = append(cheeses, "Reblchon", "Picodon")
	fmt.Println(cheeses[3])
	fmt.Println(cheeses[4])

	fmt.Println(cheeses)
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)

	var smellycheeses = make([]string, 2)
	copy(smellycheeses, cheeses)
	fmt.Println(smellycheeses)

	var newcheeses = make([]string, 3)
	copy(newcheeses, cheeses[1:])
	fmt.Println(newcheeses)

	// map
	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	fmt.Println(players["cook"])
	fmt.Println(players["stokes"])
	delete(players, "cook")
	fmt.Println(players)
}
