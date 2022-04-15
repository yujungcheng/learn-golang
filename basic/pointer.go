package main

import (
	"fmt"
)

func showMemoryAddress(x int) {
	fmt.Println(&x)
	return
}

func showMemoryAddress2(x *int) {
	fmt.Println(x)
	fmt.Println(*x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMemoryAddress(i)
	showMemoryAddress2(&i)
}
