package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "Hello World\n"
	err := ioutil.WriteFile("new_file.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
