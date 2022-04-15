/*
Go has two method to allocate memory. one is "new" and another is "make". They behavior differently and use in different cases.

the "new" method returns pointer, but "make" does not. "make" allocates array and returns a slice that referes to the array.

https://blog.wu-boy.com/2021/06/what-is-different-between-new-and-make-in-golang/
https://www.godesignpatterns.com/2014/04/new-vs-make.html
*/

package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
	foo    int
	bar    string
}

func main() {
	p := new(SyncedBuffer)
	fmt.Println("foo:", p.foo)
	fmt.Println("bar:", p.bar)
	fmt.Printf("%#v\n", p)

	fmt.Println()

	// to initialise
	p = &SyncedBuffer{
		foo: 100,
		bar: "foobar",
	}
	fmt.Println("foo:", p.foo)
	fmt.Println("bar:", p.bar)
	fmt.Printf("%#v\n", p)

	fmt.Println()

	// use make in "slice", "map" and "channel"
	np := make(map[string]string)
	np["a"] = "x"
	np["b"] = "y"
	fmt.Println(np)
}
