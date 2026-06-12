package models

type GameState struct {
	Board        Board
	InitialBoard Board
	Steps        []Step
}

func (g *GameState) Valid() error {
	return g.Board.Valid()
}
