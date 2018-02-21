package main

import (
	"fmt"
	"log"

	. "github.com/asday/pokemon-tcgo/lib"
)

func main() {
	p1 := Player{}
	p2 := Player{}

	if err := GetInput("Player 1's name:  ", &p1.Name); err != nil {
		log.Fatal(err.Error())
	}
	if err := GetInput("Player 2's name:  ", &p2.Name); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("1: %s, 2: %s\n", p1.Name, p2.Name)
}
