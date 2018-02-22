package lib

import (
	"errors"
	"math/rand"
)

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

type CardType int

const (
	TrainerCard CardType = iota
	PokemonCard
	EnergyCard
)

type Card struct {
	Name     string
	CardType CardType

	TrainerCardDetails
	PokemonCardDetails
	Type Element
}

type TrainerCardDetails struct {
	// Effect func() int ?
}

type PokemonCardDetails struct {
	EvolutionStage EvolutionStage
	EvolvesFrom    string
	Hp             int

	// PokemonPower PokemonPower
	// Moves        []Move

	WeakTo      Element
	ResistantTo Element
	Resistance  Resistance
	RetreatCost int
}

var Rattata = Card{
	Name:     "Rattata",
	CardType: PokemonCard,

	Type: Colourless,

	PokemonCardDetails: PokemonCardDetails{
		EvolutionStage: Basic,
		Hp:             30,
		WeakTo:         Rock,
		ResistantTo:    Psychic,
		Resistance:     OriginalEraResistance,
		RetreatCost:    0,
	},
}

var Raticate = Card{
	Name:     "Raticate",
	CardType: PokemonCard,

	Type: Colourless,

	PokemonCardDetails: PokemonCardDetails{
		EvolutionStage: Stage1,
		EvolvesFrom:    "Rattata",
		Hp:             60,
		WeakTo:         Rock,
		ResistantTo:    Psychic,
		Resistance:     OriginalEraResistance,
		RetreatCost:    1,
	},
}

var GrassEnergy = Card{
	Name:     "Grass Energy",
	CardType: EnergyCard,

	Type: Grass,
}

type ActivePokemon Card

type collection []Card
type Bench collection
type Deck collection
type DiscardPile collection
type Hand collection
type PrizeCards collection

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (c collection) PokemonCards() (cards []Card) {
	// TODO:  Test this.
	for _, card := range c {
		if card.CardType == PokemonCard {
			cards = append(cards, card)
		}
	}

	return
}

func (c collection) BasicPokemon() (cards []Card) {
	// TODO:  Test this.
	for _, card := range c.PokemonCards() {
		if card.PokemonCardDetails.EvolutionStage == Basic {
			cards = append(cards, card)
		}
	}

	return
}

func MoveCards(from, to []Card, amount int) error {
	// TODO:  Test this.
	if len(from) < amount {
		return errors.New("not enough cards in `from`")
	}

	for _, item := range from[:amount] {
		to = append(to, item)
	}

	from = from[amount:]

	return nil
}
