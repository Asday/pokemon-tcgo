package lib

import (
	"errors"
	"fmt"
	"html/template"
)

type gamePhase int

const (
	setUp gamePhase = iota
	play
)

type Game struct {
	players []Player

	decks         []Deck
	hands         []Hand
	discardPiles  []DiscardPile
	benches       []Bench
	activePokemon []ActivePokemon
	prizeCards    []PrizeCards

	currentPlayer  int
	attachedEnergy bool
	phase          gamePhase
}

func NewGame(players []Player, decks []Deck, firstPlayer int) Game {
	game := Game{
		players: players,
		decks:   decks,

		hands:         make([]Hand, len(players)),
		discardPiles:  make([]DiscardPile, len(players)),
		benches:       make([]Bench, len(players)),
		activePokemon: make([]ActivePokemon, len(players)),
		prizeCards:    make([]PrizeCards, len(players)),

		currentPlayer: firstPlayer,
		phase:         setUp,
	}

	for _, deck := range game.decks {
		deck.Shuffle()
	}

	for player := range game.players {
		game.Draw(player, 7)
	}

	return game
}

func (g *Game) AdvanceTurn() {
	// TODO:  Test this.
	g.currentPlayer++
	if g.currentPlayer >= len(g.players) {
		g.currentPlayer = 0
	}
}

func (g *Game) AdvancePhase() error {
	switch g.phase {
	case setUp:
		// TODO:  Prize cards.
		g.phase = play
		return nil
	}

	return errors.New("no next phase")
}

func (g *Game) Draw(player int, cards int) {
	MoveCards(g.decks[player], g.hands[player], cards)
}

type Action func()

type ActionInfo struct {
	Prompt *template.Template
	Player int
	Action Action
}

func (g *Game) GetActions() (actions []ActionInfo) {
	switch g.phase {
	case setUp:
		// Mulligans.
		for player, hand := range g.hands {
			if len(Collection(hand).BasicPokemon()) == 0 {
				actions = append(actions, ActionInfo{
					Prompt: template.Must(
						template.New("").Parse("{{.Name}} has no basic Pokémon!"),
					),
					Player: player,
					Action: func() {
						g.Mulligan(player)
					},
				})
			}
		}
	}

	return
}

func (g *Game) Mulligan(player int) {
	fmt.Println("mulligan")
}
