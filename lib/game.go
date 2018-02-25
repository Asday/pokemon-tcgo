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

	placedBenchedPokemon []bool
	alive                []bool
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

		placedBenchedPokemon: make([]bool, len(players)),
		alive:                make([]bool, len(players)),
	}

	for player := range players {
		game.hands[player] = make(Hand, 0)
		game.discardPiles[player] = make(DiscardPile, 0)
		game.activePokemon[player] = make(ActivePokemon, 0)
		game.benches[player] = make(Bench, 0)
		game.prizeCards[player] = make(PrizeCards, 0)

		game.placedBenchedPokemon[player] = false
		game.alive[player] = true
	}

	for _, deck := range game.decks {
		deck.Shuffle()
	}

	for player := range game.players {
		game.Draw(player, 7)
	}

	return game
}

func (g *Game) advanceCurrentPlayer() {
	// TODO:  Test this.
	g.currentPlayer++
	if g.currentPlayer >= len(g.players) {
		g.currentPlayer = 0
	}
}

func (g *Game) AdvanceTurn() {
	g.advanceCurrentPlayer()

	g.attachedEnergy = false
	g.Draw(g.currentPlayer, 1)
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
		g.Draw(g.currentPlayer, 1)
		g.phase = play
		return nil
	}

	return errors.New("no next phase")
}

func (g *Game) Draw(player int, cards int) error {
	deck, hand, err := MoveCards(
		Collection(g.decks[player]),
		Collection(g.hands[player]),
		cards,
	)
	if err != nil {
		return err
	}

	g.decks[player] = Deck(deck)
	g.hands[player] = Hand(hand)

	if cards == 1 {
		fmt.Printf("%s draws a card.\n", g.players[player].Name)
	} else {
		fmt.Printf("%s draws %d cards.\n", g.players[player].Name, cards)
	}

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

type Actions []ActionInfo

func (a Actions) Choices() (choices []string) {
	for _, action := range a {
		choices = append(choices, action.Prompt)
	}

	return
}

func (g *Game) GetActions() (actions Actions) {
	switch g.phase {
	case mulligans:
		for player, hand := range g.hands {
			if len(Collection(hand).BasicPokemon()) == 0 {
				playerIndex := player
				actions = append(actions, ActionInfo{
					Prompt: fmt.Sprintf(
						"%s has no basic Pokémon!",
						g.players[playerIndex].Name,
					),
					Action: func() {
						g.Mulligan(playerIndex)
					},
				})

				return
			}
		}

		return

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

				return
			}
		}

		return

	case placeBenchedPokemon:
		for i, player := range g.players {
			if !g.placedBenchedPokemon[i] {
				playerIndex := i
				actions = append(actions, ActionInfo{
					Prompt: fmt.Sprintf(
						"%s, choose up to 5 basic Pokémon to place on the bench.",
						player.Name,
					),
					Action: func() {
						g.PlaceBenchedPokemon(playerIndex)
					},
				})

				return
			}
		}

		return

	case play:
		if g.GameOver() {
			return
		}
	}

	panic(fmt.Sprintf("unhandled phase:  %v", g.phase))
}

func (g Game) GameOver() bool {
	playersAlive := 0

	for playerIndex := range g.players {
		if g.alive[playerIndex] {
			playersAlive += 1
		}
	}

	return playersAlive <= 1
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

	fmt.Printf(
		"%s shuffles their hand into their deck.\n",
		g.players[player].Name,
	)
}

func (g *Game) Mulligan(player int) {
	g.RevealHand(player)
	g.ShuffleHandIntoDeck(player)
	g.Draw(player, 7)
}

func (g Game) RevealHand(player int) {
	fmt.Printf("%s shows their hand.\n", g.players[player].Name)
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

func (g *Game) PlaceBenchedPokemon(player int) {
	// TODO:  Implement.
	indices := Collection(g.hands[player]).GetCardChoices(BasicPokemonValidator, 5)

	hand, bench := MoveCardsAtIndices(
		Collection(g.hands[player]),
		Collection(g.benches[player]),
		indices,
	)

	g.hands[player] = Hand(hand)
	g.benches[player] = Bench(bench)

	g.placedBenchedPokemon[player] = true
}
