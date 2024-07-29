package bot

import (
	"math/rand"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/match"
)

func GetMove(state *match.State) *board.Move {

	moves := state.Board.MovesForColor(state.Board.Turn, false)
	var consideredMoves []*board.Move
	var consideredMoveEval int

	for _, move := range moves {
		b := state.Board.DeepCopy()
		b.ExecuteMove(move)
		eval := b.Evaluate()
		if eval > consideredMoveEval {
			// only this one move in array
			consideredMoves = append([]*board.Move{move}, nil)
			consideredMoveEval = eval
		} else if eval == consideredMoveEval {
			// add to array
			consideredMoves = append(consideredMoves, move)
		}
	}

	// chose one of the best Moves
	return consideredMoves[rand.Intn(len(consideredMoves))] //at random positon
}
