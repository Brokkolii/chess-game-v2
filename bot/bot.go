package bot

import (
	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/match"
)

func GetMove(state *match.State) *board.Move {
	// First select a random Move
	moves := state.Board.MovesForColor(state.Turn.Color, false)
	var consideredMove *board.Move
	var consideredMoveEval int

	for _, move := range moves {
		b := state.Board.DeepCopy()
		b.ExecuteMove(move)
		// TODO: doesnt work, cause we need the full state!
		// => for this we need to be able to deepclone the whole state
		// => make it so we can transfer state to fen and back
		// => alot of the stuff that board does needs to be done by state
		// => calculate the future with future states and not boards
		s := match.NewState()
		s.Board = b
		eval := s.Evaluate()
		if eval >= consideredMoveEval {
			consideredMove = move
			consideredMoveEval = eval
		}
	}

	return consideredMove
}
