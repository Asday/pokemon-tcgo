package main

import (
	"fmt"
	"log"

	. "github.com/asday/pokemon-tcgo/lib"
)

func main() {
	players := make([]Player, 2)
	players = append(players, Player{})
	players = append(players, Player{})

	if err := GetInput("Player 1's name:  ", &players[0].Name); err != nil {
		log.Fatal(err.Error())
	}
	if err := GetInput("Player 2's name:  ", &players[1].Name); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("1: %s, 2: %s\n", players[0].Name, players[1].Name)
}
