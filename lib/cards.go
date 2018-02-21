package lib

import "math/rand"

// This file is likely to split in future.

type EvolutionStage int

const (
	Basic EvolutionStage = iota
	Stage1
	Stage2
)

type Resistance func(damage int) int

func OriginalEraResistance(damage int) int {
	damage -= 30

	if damage > 0 {
		return damage
	}

	return 0
}

type Card struct {
	Name string
}

type TrainerCard struct {
	Card
}

type PokemonCard struct {
	Card

	EvolutionStage EvolutionStage
	EvolvesFrom    string
	Hp             int
	Type           Element

	// PokemonPower PokemonPower
	// Moves        []Move

	WeakTo      Element
	ResistantTo Element
	Resistance  Resistance
	RetreatCost int
}

var Ratatta = PokemonCard{
	Card: Card{
		Name: "Rattata",
	},
	EvolutionStage: Basic,
	Hp:             30,
	Type:           Colourless,
	WeakTo:         Rock,
	ResistantTo:    Psychic,
	Resistance:     OriginalEraResistance,
	RetreatCost:    0,
}

var Raticate = PokemonCard{
	Card: Card{
		Name: "Raticate",
	},
	EvolutionStage: Stage1,
	EvolvesFrom:    "Rattata",
	Hp:             60,
	Type:           Colourless,
	WeakTo:         Rock,
	ResistantTo:    Psychic,
	Resistance:     OriginalEraResistance,
	RetreatCost:    1,
}

type EnergyCard struct {
	Card

	Type Element
}

var GrassEnergy = EnergyCard{
	Card: Card{
		Name: "Grass Energy",
	},
	Type: Grass,
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
