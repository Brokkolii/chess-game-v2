package board

func (b *Board) Evaluate() int {
	evaluation := 0
	evaluation += b.evaluateByPieceValue()
	return evaluation
}

func (b *Board) evaluateByPieceValue() int {
	evaluation := 0
	for _, piece := range b.Pieces() {
		if piece != nil {
			pieceValue := 0
			switch piece.Type {
			case "pawn":
				pieceValue = 1
			case "rook":
				pieceValue = 5
			case "knight":
				pieceValue = 3
			case "bishop":
				pieceValue = 3
			case "queen":
				pieceValue = 9
			case "king":
				pieceValue = 10
			}
			if piece.Color == b.Turn {
				evaluation += pieceValue
			} else {
				evaluation += pieceValue * -1
			}
		}
	}
	return evaluation
}
