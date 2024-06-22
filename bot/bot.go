package bot

import (
	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/match"
)

func GetMove(state *match.State) *board.Move {

	moves := state.Board.MovesForColor(state.Board.Turn, false)
	var consideredMove *board.Move
	var consideredMoveEval int

	for _, move := range moves {
		b := state.Board.DeepCopy()
		b.ExecuteMove(move)
		eval := b.Evaluate()
		if eval >= consideredMoveEval {
			consideredMove = move
			consideredMoveEval = eval
		}
	}

	return consideredMove
}
