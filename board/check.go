package board

import "github.com/Brokkolii/chess-game-v2/util"

func (b *Board) isInCheckAfterMove(move Move) bool {
	color := move.From.Piece.Color
	newBoard := b.deepCopy()
	newBoard.ExecuteMove(&move)
	return newBoard.isInCheck(color)
}

func (b *Board) IsCheckMate(color string) bool {
	// TODO Method is not working yet :D
	if b.isInCheck(color) {
		moves := b.MovesForColor(color, false)
		if len(moves) == 0 {
			return true
		}
	}
	return false
}

func (b *Board) isInCheck(color string) bool {
	moves := b.MovesForColor(util.InvertColor(color), true)
	for _, move := range moves {
		if move.To.Piece != nil && move.To.Piece.Type == "king" && move.To.Piece.Color == color {
			//fmt.Println(util.InvertColor(color), move.From.Piece.Type, move.From.Row, move.From.Col, "attacking", color, "king")
			return true
		}
	}

	return false
}
