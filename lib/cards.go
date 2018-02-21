package lib

import "math/rand"

// This file is likely to split in future.

type EvolutionStage int

const (
	Basic EvolutionStage = iota
	Stage1
	Stage2
)

type Card struct {
	Name string
}

type TrainerCard struct {
	Card
}

type PokemonCard struct {
	Card

	EvolutionStage EvolutionStage
	Hp             int
	Type           Element

	// PokemonPower PokemonPower
	// Moves        []Move

	WeakTo      Element
	ResistantTo Element
	RetreatCost int
}

type EnergyCard struct {
	Card

	Type Element
}

type ActivePokemon Card
type Bench []Card
type Deck []Card
type DiscardPile []Card
type Hand []Card
type PrizeCards []Card

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
