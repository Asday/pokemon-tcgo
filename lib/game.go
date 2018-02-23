package lib

import (
	"errors"
	"fmt"
)

type gamePhase int

const (
	mulligans gamePhase = iota
	placeActivePokemon
	placeBenchedPokemon
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
		phase:         mulligans,
	}

	for player := range players {
		game.hands[player] = make(Hand, 0)
		game.discardPiles[player] = make(DiscardPile, 0)
		game.activePokemon[player] = make(ActivePokemon, 0)
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
	case mulligans:
		g.phase = placeActivePokemon
		return nil
	case placeActivePokemon:
		g.phase = placeBenchedPokemon
		return nil
	case placeBenchedPokemon:
		g.placePrizeCards(6)
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

func (g *Game) placePrizeCards(cards int) error {
	for player := range g.players {
		deck, prizeCards, err := MoveCards(
			Collection(g.decks[player]),
			Collection(g.prizeCards[player]),
			cards,
		)
		if err != nil {
			return err
		}

		g.decks[player] = Deck(deck)
		g.prizeCards[player] = PrizeCards(prizeCards)
	}

	return nil
}

type Action func()

type ActionInfo struct {
	Prompt string
	Action Action
}

func (g *Game) GetActions() (actions []ActionInfo) {
	switch g.phase {
	case mulligans:
		for player, hand := range g.hands {
			if len(Collection(hand).BasicPokemon()) == 0 {
				playerIndex := player
				actions = append(actions, ActionInfo{
					Prompt: fmt.Sprintf(
						"%s has no basic Pokémon!\n\n%s shows their hand.",
						g.players[playerIndex].Name,
						g.players[playerIndex].Name,
					),
					Action: func() {
						g.Mulligan(playerIndex)
					},
				})
			}
		}
	case placeActivePokemon:
		for i, player := range g.players {
			if len(g.activePokemon[i]) == 0 {
				playerIndex := i
				actions = append(actions, ActionInfo{
					Prompt: fmt.Sprintf(
						"%s, choose an active Pokémon.",
						player.Name,
					),
					Action: func() {
						g.PlaceActivePokemon(playerIndex)
					},
				})
			}
		}
	case placeBenchedPokemon:
		fmt.Println("place benched")
	case play:
		fmt.Println("play")
	}

	return
}

func (g *Game) ShuffleHandIntoDeck(player int) {
	// TODO: Test this.
	// Wonder what happens when you move an empty hand to an empty deck.
	hand, deck, _ := MoveCards( // Error doesn't matter.
		Collection(g.hands[player]),
		Collection(g.decks[player]),
		len(g.hands[player]),
	)

	g.hands[player] = Hand(hand)
	g.decks[player] = Deck(deck)

	g.decks[player].Shuffle()
}

func (g *Game) Mulligan(player int) {
	g.RevealHand(player)

	fmt.Printf(
		"%s shuffles their hand into their deck and draws a new one.\n\n",
		g.players[player],
	)
	g.ShuffleHandIntoDeck(player)
	g.Draw(player, 7)
}

func (g Game) RevealHand(player int) {
	for _, card := range g.hands[player] {
		fmt.Println(card)
	}

	Next()
}

func (g *Game) PlaceActivePokemon(player int) {
	index := Collection(g.hands[player]).GetCardChoice(BasicPokemonValidator)

	hand, activePokemon := MoveCardAtIndex(
		Collection(g.hands[player]),
		Collection(g.activePokemon[player]),
		index,
	)

	g.hands[player] = Hand(hand)
	g.activePokemon[player] = ActivePokemon(activePokemon)
}
