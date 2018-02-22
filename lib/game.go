package lib

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
		players:       players,
		decks:         decks,
		currentPlayer: firstPlayer,
		phase:         setUp,
	}

	for _, deck := range game.decks {
		deck.Shuffle()
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
