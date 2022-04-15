package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

// Omitting empty struct
type Person2 struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

func main() {
	hobbies := []string{"Cycling", "Cheese", "Technology"}
	p := Person{
		Name:    "Tom",
		Age:     36,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)

	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)

	// empty struct
	p2 := Person2{}
	fmt.Printf("%+v\n", p2)

	jsonByteData2, err := json.Marshal(p2)
	if err != nil {
		log.Fatal(err)
	}

	jsonStringData2 := string(jsonByteData2)
	fmt.Println(jsonStringData2)
}
