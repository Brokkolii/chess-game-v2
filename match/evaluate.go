package match

func (s *State) Evaluate() int {
	evaluation := 0
	evaluation += s.evaluateByPieceValue()
	return evaluation
}

func (s *State) evaluateByPieceValue() int {
	evaluation := 0
	for _, piece := range s.Board.Pieces() {
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
			if piece.Color == s.Turn.Color {
				evaluation += pieceValue
			} else {
				evaluation += pieceValue * -1
			}
		}
	}
	return evaluation
}
