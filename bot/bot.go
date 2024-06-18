package bot

import (
	"math/rand"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/game"
)

func PlayMove(state *game.GameState) {
	move := ChooseMove(state)
	state.Board.ExecuteMove(move)
	state.NextTurn()
}

func ChooseMove(state *game.GameState) *board.Move {
	moves := state.Board.MovesForColor(state.Turn).Moves
	randomIndex := rand.Intn(len(moves))
	return moves[randomIndex]
}
