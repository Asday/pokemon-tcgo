package lib

import (
	"errors"
	"fmt"
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

func (c Card) String() string {
	return c.Name
}

func (c Card) IsNull() bool {
	return c.Name == ""
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

type Collection []Card
type ActivePokemon Collection
type Bench Collection
type Deck Collection
type DiscardPile Collection
type Hand Collection
type PrizeCards Collection

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (c Collection) PokemonCards() (cards Collection) {
	// TODO:  Test this.
	for _, card := range c {
		if card.CardType == PokemonCard {
			cards = append(cards, card)
		}
	}

	return
}

func (c Collection) BasicPokemon() (cards Collection) {
	// TODO:  Test this.
	for _, card := range c.PokemonCards() {
		if card.PokemonCardDetails.EvolutionStage == Basic {
			cards = append(cards, card)
		}
	}

	return
}

type CardValidator func(card Card) (bool, string)

func BasicPokemonValidator(card Card) (bool, string) {
	if card.CardType != PokemonCard {
		return false, "That wasn't a Pokémon card."
	}

	if card.PokemonCardDetails.EvolutionStage != Basic {
		return false, "That wasn't a Basic Pokémon"
	}

	return true, ""
}

func (c Collection) GetCardChoice(validator CardValidator) int {
	var choices []string
	for _, card := range c {
		choices = append(choices, card.String())
	}

	var choiceIndex int
	for {
		choiceIndex = GetChoice("", choices)

		if valid, message := validator(c[choiceIndex]); valid {
			break
		} else {
			fmt.Println(message)
		}
	}

	return choiceIndex
}

func MoveCards(from, to Collection, amount int) (Collection, Collection, error) {
	// TODO:  Test this.
	if len(from) < amount {
		return nil, nil, errors.New("not enough cards in `from`")
	}

	for _, item := range from[:amount] {
		to = append(to, item)
	}

	from = from[amount:]

	return from, to, nil
}

func MoveCardsAtIndices(from, to Collection, indices []int) (Collection, Collection) {
	// TODO:  Test this.
	indexSet := make(map[int]struct{})
	for _, index := range indices {
		to = append(to, from[index])
		indexSet[index] = struct{}{}
	}

	newFrom := make(Collection, 0)
	for i, value := range from {
		if _, ok := indexSet[i]; !ok {
			newFrom = append(newFrom, value)
		}
	}

	return newFrom, to
}

func MoveCardAtIndex(from, to Collection, index int) (Collection, Collection) {
	return MoveCardsAtIndices(from, to, []int{index})
}
