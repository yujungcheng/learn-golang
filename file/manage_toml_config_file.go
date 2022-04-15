package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Name   string `json:"name"`
	Awake  bool   `json:"awake"`
	Hungry bool   `json:"hungry"`
}

func main() {
	c := Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)
}
