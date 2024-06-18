package game

import (
	"github.com/Brokkolii/chess-game-v2/board"
)

type GameState struct {
	Board *board.Board
	Turn  string
}

func NewGameState() *GameState {
	return &GameState{
		Board: board.NewBoard(),
		Turn:  "white",
	}
}

func (gs *GameState) NextTurn() {
	gs.Turn = InvertColor(gs.Turn)
}

func InvertColor(color string) string {
	if color == "white" {
		return "black"
	} else {
		return "white"
	}
}
