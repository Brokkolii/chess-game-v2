package bot

import (
	"math/rand"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/match"
)

func GetMove(state *match.State) *board.Move {
	move := ChooseMove(state)
	return move
}

func ChooseMove(state *match.State) *board.Move {
	moves := state.Board.MovesForColor(state.Turn.Color).Moves
	randomIndex := rand.Intn(len(moves))
	return moves[randomIndex]
}
