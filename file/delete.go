package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	deleteFile := "./old_file.txt"
	err := os.Remove(deleteFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v deleted\n", deleteFile)
}
