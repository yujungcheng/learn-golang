package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fromFile := "./new_file.txt"
	toFile := "old_file.txt"

	// method1: io.Copy
	from, err := os.Open(fromFile)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close() // close file after all other execution finished

	to, err := os.OpenFile(toFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close() // close file after all other execution finished

	fmt.Printf("Copy files from %v to %v \n", fromFile, toFile)
	_, err = io.Copy(to, from) // do copy
	if err != nil {
		log.Fatal(err)
	}

	// method2: ioutil.WriteFile() and ioutil.ReadFile()
	input, err := ioutil.ReadFile(fromFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	toFile = "old_file2.txt"
	err = ioutil.WriteFile(toFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", toFile)
		fmt.Println(err)
		return
	}
	fmt.Printf("Copy files from %v to %v \n", fromFile, toFile)

	// method3: os.Read() and os.Write()
	// https://github.com/mactsouk/opensource.com/blob/master/cp3.go
	/*
		buf := make([]byte, BUFFERSIZE)
		for {
				n, err := source.Read(buf)
				if err != nil && err != io.EOF {
						return err
				}
				if n == 0 {
						break
				}

				if _, err := destination.Write(buf[:n]); err != nil {
						return err
				}
		}
	*/
}
