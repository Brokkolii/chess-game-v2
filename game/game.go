package game

import (
	"github.com/Brokkolii/chess-game-v2/board"
)

type GameState struct {
	Board *board.Board
	turn  string
}

func NewGameState() *GameState {
	return &GameState{
		Board: board.NewBoard(),
		turn:  "white",
	}
}
