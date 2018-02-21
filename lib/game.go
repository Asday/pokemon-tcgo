package lib

type Game struct {
	players []Player

	decks         []Deck
	hands         []Hand
	discardPiles  []DiscardPile
	benches       []Bench
	activePokemon []ActivePokemon

	currentPlayer  int
	attachedEnergy bool
}

func NewGame(players []Player, firstPlayer int) Game {
	return Game{
		players:       players,
		currentPlayer: firstPlayer,
	}
}

func (g *Game) AdvanceTurn() {
	// TODO:  Test this.
	g.currentPlayer++
	if g.currentPlayer >= len(g.players) {
		g.currentPlayer = 0
	}
}
