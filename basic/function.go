package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func getPrize() (int, string) {
	i := 2
	s := "gold fish"
	return i, s
}

func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func sayHi() (x, y string) {
	x = "Hello"
	y = "World"
	return
}

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full!, I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)
	return feedMe(portion, eaten)
}

func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fmt.Printf("%v\n", isEven(1))
	fmt.Printf("%v\n", isEven(2))

	quantity, prize := getPrize()
	fmt.Printf("You won %v %v\n", quantity, prize)

	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("The result is %v\n", result)

	fmt.Println(sayHi())

	fmt.Println(feedMe(1, 0))

	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFunction(fn))
}
