package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"os"

	. "github.com/asday/pokemon-tcgo/lib"
)

func main() {
	players := make([]Player, 2)
	decks := make([]Deck, 2)

	decks[0] = Deck{Rattata, Raticate, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy}
	decks[1] = Deck{Rattata, Rattata, Rattata, Rattata, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy, GrassEnergy}

	if err := GetInput("Player 1's name:  ", &players[0].Name); err != nil {
		log.Fatal(err.Error())
	}
	if err := GetInput("Player 2's name:  ", &players[1].Name); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s vs %s!\n\n", players[0].Name, players[1].Name)
	fmt.Print("Flipping coin to see who goes first")

	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
	}

	time.Sleep(1 * time.Second)

	fmt.Println()

	firstPlayer := rand.Intn(2)

	fmt.Printf("%s wins the toss and goes first!\n", players[firstPlayer].Name)

	game := NewGame(players, decks, firstPlayer)

	for {
		actions := game.GetActions()

		if len(actions) == 0 {
			if err := game.AdvancePhase(); err != nil {
				break
			}
		}

		for _, action := range actions {
			action.Prompt.Execute(os.Stdout, players[action.Player])
			fmt.Println()
			action.Action()
		}

		break // TODO:  Remove.
	}
}
