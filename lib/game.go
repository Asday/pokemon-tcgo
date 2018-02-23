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

	for player := range players {
		game.hands[player] = make(Hand, 0)
		game.discardPiles[player] = make(DiscardPile, 0)
		game.benches[player] = make(Bench, 0)
		game.prizeCards[player] = make(PrizeCards, 0)
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

func (g *Game) Draw(player int, cards int) error {
	deck, hand, err := MoveCards(Collection(g.decks[player]), Collection(g.hands[player]), cards)
	if err != nil {
		return err
	}

	g.decks[player] = Deck(deck)
	g.hands[player] = Hand(hand)

	return nil
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
				playerIndex := player
				actions = append(actions, ActionInfo{
					Prompt: template.Must(
						template.New("").Parse("{{.Name}} has no basic Pokémon!\n\n{{.Name}} shows their hand."),
					),
					Player: player,
					Action: func() {
						g.Mulligan(playerIndex)
					},
				})
			}
		}
	}

	return
}

func (g *Game) Mulligan(player int) {
	g.RevealHand(player)
}

func (g Game) RevealHand(player int) {
	for _, card := range g.hands[player] {
		fmt.Println(card)
	}
}
