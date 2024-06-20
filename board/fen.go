package board

import "fmt"

func (b *Board) toFEN() string {
	fen := ""
	epmtyFieldsCounter := 0
	for row := 8; row >= 1; row-- {
		for col := 1; col <= 8; col++ {
			piece := b.FieldAt(row, col).Piece
			if piece != nil {
				if epmtyFieldsCounter > 0 {
					fen = fen + fmt.Sprint(epmtyFieldsCounter)
				}
				epmtyFieldsCounter = 0
				// ad letter
				if piece.Color == "white" {
					switch piece.Type {
					case "rook":
						fen = fen + "R"
					case "bishop":
						fen = fen + "B"
					case "knight":
						fen = fen + "N"
					case "queen":
						fen = fen + "Q"
					case "king":
						fen = fen + "K"
					case "pawn":
						fen = fen + "P"
					}
				} else {
					switch piece.Type {
					case "rook":
						fen = fen + "r"
					case "bishop":
						fen = fen + "b"
					case "knight":
						fen = fen + "n"
					case "queen":
						fen = fen + "q"
					case "king":
						fen = fen + "k"
					case "pawn":
						fen = fen + "p"
					}
				}
			} else {
				epmtyFieldsCounter++
			}
		}
		if epmtyFieldsCounter > 0 {
			fen = fen + fmt.Sprint(epmtyFieldsCounter)
		}
		epmtyFieldsCounter = 0
		fen = fen + "/"
	}

	return fen + " "
}

func (b *Board) fromFEN(fen string) {
	row := 8
	col := 1
	for _, char := range fen {
		switch char {
		case 'r':
			piece := NewPiece("black", "rook", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'n':
			piece := NewPiece("black", "knight", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'b':
			piece := NewPiece("black", "bishop", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'q':
			piece := NewPiece("black", "queen", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'k':
			piece := NewPiece("black", "king", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'p':
			piece := NewPiece("black", "pawn", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'R':
			piece := NewPiece("white", "rook", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'N':
			piece := NewPiece("white", "knight", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'B':
			piece := NewPiece("white", "bishop", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'Q':
			piece := NewPiece("white", "queen", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'K':
			piece := NewPiece("white", "king", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'P':
			piece := NewPiece("white", "pawn", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case '/':
			row--
			col = 1
		case ' ':
			return
		default:
			if char >= '1' && char <= '8' {
				for i := 1; i <= int(char-'0'); i++ {
					field := NewField(row, col)
					b.AddField(field)
					col++
				}
			}
		}
	}
}
