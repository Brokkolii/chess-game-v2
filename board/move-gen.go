package board

func (b *Board) MovesForPiece(piece *Piece, ignorePins bool) []*Move {
	var moves []*Move
	// TODO: aun pasante
	if piece.Type == "pawn" {
		if piece.Color == "white" {
			// normal forward
			row := piece.Field.Row + 1
			col := piece.Field.Col
			forwardField := b.FieldAt(row, col)
			if forwardField != nil && forwardField.Piece == nil {
				move := NewMove(piece.Field, forwardField)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}

				// double forward
				if piece.Field.Row == 2 {
					row := piece.Field.Row + 2
					col := piece.Field.Col
					doubleForwardField := b.FieldAt(row, col)
					if doubleForwardField != nil && doubleForwardField.Piece == nil {
						move := NewMove(piece.Field, doubleForwardField)
						if ignorePins || !b.isInCheckAfterMove(*move) {
							moves = append(moves, move)
						}
					}
				}
			}

			// takes
			leftTakeField := b.FieldAt(piece.Field.Row+1, piece.Field.Col-1)
			if leftTakeField != nil && leftTakeField.Piece != nil && leftTakeField.Piece.Color != piece.Color {
				move := NewMove(piece.Field, leftTakeField)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			}
			rightTakeField := b.FieldAt(piece.Field.Row+1, piece.Field.Col+1)
			if rightTakeField != nil && rightTakeField.Piece != nil && rightTakeField.Piece.Color != piece.Color {
				move := NewMove(piece.Field, rightTakeField)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			}

		} else {
			// normal forward
			row := piece.Field.Row - 1
			col := piece.Field.Col
			forwardField := b.FieldAt(row, col)
			if forwardField != nil && forwardField.Piece == nil {
				move := NewMove(piece.Field, forwardField)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}

				// double forward
				if piece.Field.Row == 7 {
					row := piece.Field.Row - 2
					col := piece.Field.Col
					doubleForwardField := b.FieldAt(row, col)
					if doubleForwardField != nil && doubleForwardField.Piece == nil {
						move := NewMove(piece.Field, doubleForwardField)
						if ignorePins || !b.isInCheckAfterMove(*move) {
							moves = append(moves, move)
						}
					}
				}
			}

			// takes
			leftTakeField := b.FieldAt(piece.Field.Row-1, piece.Field.Col-1)
			if leftTakeField != nil && leftTakeField.Piece != nil && leftTakeField.Piece.Color != piece.Color {
				move := NewMove(piece.Field, leftTakeField)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			}
			rightTakeField := b.FieldAt(piece.Field.Row-1, piece.Field.Col+1)
			if rightTakeField != nil && rightTakeField.Piece != nil && rightTakeField.Piece.Color != piece.Color {
				move := NewMove(piece.Field, rightTakeField)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			}
		}
	}
	if piece.Type == "rook" || piece.Type == "queen" {
		// upwards
		for i := piece.Field.Row + 1; i <= size; i++ {
			field := b.FieldAt(i, piece.Field.Col)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
		// downwards
		for i := piece.Field.Row - 1; i >= 1; i-- {
			field := b.FieldAt(i, piece.Field.Col)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
		// left
		for i := piece.Field.Col - 1; i >= 1; i-- {
			field := b.FieldAt(piece.Field.Row, i)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
		// right
		for i := piece.Field.Col + 1; i <= size; i++ {
			field := b.FieldAt(piece.Field.Row, i)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
	}

	if piece.Type == "bishop" || piece.Type == "queen" {
		// top right
		for i := 1; i <= size; i++ {
			field := b.FieldAt(piece.Field.Row+i, piece.Field.Col+i)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
		// bottom right
		for i := 1; i <= size; i++ {
			field := b.FieldAt(piece.Field.Row-i, piece.Field.Col+i)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
		// bottom left
		for i := 1; i <= size; i++ {
			field := b.FieldAt(piece.Field.Row-i, piece.Field.Col-i)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
		// top left
		for i := 1; i <= size; i++ {
			field := b.FieldAt(piece.Field.Row+i, piece.Field.Col-i)

			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {
					moves = append(moves, move)
				}
				break
			} else {
				break
			}
		}
	}

	if piece.Type == "knight" {
		// maybe i can use this offset technique on all moves
		offsets := [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
		for _, offset := range offsets {
			field := b.FieldAt(piece.Field.Row+offset[0], piece.Field.Col+offset[1])
			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {

					moves = append(moves, move)

				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {

					moves = append(moves, move)

				}
			}
		}
	}

	if piece.Type == "king" {
		offsets := [][]int{{1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}}
		for _, offset := range offsets {
			field := b.FieldAt(piece.Field.Row+offset[0], piece.Field.Col+offset[1])
			if field != nil && field.Piece == nil {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {

					moves = append(moves, move)

				}
			} else if field != nil && field.Piece != nil && field.Piece.Color != piece.Color {
				move := NewMove(piece.Field, field)
				if ignorePins || !b.isInCheckAfterMove(*move) {

					moves = append(moves, move)

				}
			}
		}
	}

	return moves
}

func (b *Board) MovesForColor(color string, ignorePins bool) []*Move {
	var moves []*Move

	for _, piece := range b.Pieces() {
		if piece.Color == color {
			moves = append(moves, b.MovesForPiece(piece, ignorePins)...)
		}
	}

	return moves
}
