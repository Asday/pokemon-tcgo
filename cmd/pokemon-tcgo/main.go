package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	. "github.com/asday/pokemon-tcgo/lib"
)

func main() {
	players := make([]Player, 2)

	if err := GetInput("Player 1's name:  ", &players[0].Name); err != nil {
		log.Fatal(err.Error())
	}
	if err := GetInput("Player 2's name:  ", &players[1].Name); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s vs %s!\n", players[0].Name, players[1].Name)
	fmt.Print("Flipping coin to see who goes first")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}

	time.Sleep(1 * time.Second)

	fmt.Println()

	firstPlayer := rand.Intn(2)

	fmt.Printf("%s wins the toss and goes first!\n", players[firstPlayer].Name)
}
