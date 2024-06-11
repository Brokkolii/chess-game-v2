package game

import "github.com/Brokkolii/chess-game-v2/board"

type Move struct {
	From *board.Field
	To   *board.Field
}
