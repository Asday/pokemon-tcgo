package lib

// This file is likely to split in future.

type Card struct {
	Name string
}

type TrainerCard struct {
	Card
}

type PokemonCard struct {
	Card

	Hp int
	// Type Element
	//
	// PokemonPower PokemonPower
	// Moves        []Move
	//
	// WeakTo      Element
	// ResistantTo Element
	RetreatCost int
}

type ActivePokemon Card
type Bench []Card
type Deck []Card
type DiscardPile []Card
type Hand []Card
