package lib

type Game struct {
	currentPlayer int
	players       []Player
}

func NewGame(players []Player, firstPlayer int) Game {
	return Game{
		firstPlayer,
		players,
	}
}

func (g *Game) AdvanceTurn() {
	// TODO:  Test this.
	g.currentPlayer++
	if g.currentPlayer >= len(g.players) {
		g.currentPlayer = 0
	}
}
