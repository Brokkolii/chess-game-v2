package board

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Fields   [size][size]*Field
	Image    *ebiten.Image
	Started  bool
	GameOver string
}

const (
	size       = 8
	squareSize = 80
)

func (b *Board) Draw(dst *ebiten.Image) {
	if b.Image == nil {
		b.Image = ebiten.NewImage(squareSize*size, squareSize*size)
		for i, row := range b.Fields {
			for j := range row {
				field := b.FieldAt(i+1, j+1)
				if field != nil {
					field.draw(b.Image)
				}
			}
		}
	}
	dst.DrawImage(b.Image, nil)
}

func (b *Board) DrawPieces(dst *ebiten.Image) {
	for _, piece := range b.Pieces() {
		piece.Draw(dst)
	}
}

func (b *Board) DrawMoves(dst *ebiten.Image, moves []*Move) {
	for _, move := range moves {
		move.Draw(dst)
	}
}

func NewBoard() *Board {
	var board Board
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board.fromFEN(fen)
	return &board
}

func NewBoardFromFEN(fen string) *Board {
	var board Board
	board.fromFEN(fen)
	return &board
}

func (b *Board) deepCopy() *Board {
	fen := b.toFEN()
	board := NewBoardFromFEN(fen)
	return board
}

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

func (b *Board) Pieces() []*Piece {
	var pieces []*Piece
	for _, row := range b.Fields {
		for _, field := range row {
			piece := field.Piece
			if piece != nil {
				pieces = append(pieces, piece)
			}
		}
	}
	return pieces
}

func (b *Board) PieceAtCoords(x int, y int) *Piece {
	for _, piece := range b.Pieces() {
		left, top := piece.getCoords()
		right := left + squareSize
		bottom := top + squareSize
		if y <= bottom && y >= top && x <= right && x >= left {
			return piece
		}
	}
	return nil
}

func (b *Board) FieldAtCoords(x int, y int) *Field {
	col := (x / squareSize) + 1
	row := (size - 1 - (y / squareSize)) + 1
	if row > 8 || row < 1 || col > 8 || col < 1 || y < 0 || x < 0 {
		return nil
	}
	return b.FieldAt(row, col)
}

func (b *Board) FieldAt(row int, col int) *Field {
	if row > 8 || row < 1 || col > 8 || col < 1 {
		return nil
	}
	return b.Fields[row-1][col-1]
}

func (b *Board) AddField(field *Field) {
	b.Fields[field.Row-1][field.Col-1] = field
}

func (b *Board) ExecuteMove(move *Move) {
	fromField := b.FieldAt(move.From.Row, move.From.Col)
	toField := b.FieldAt(move.To.Row, move.To.Col)
	piece := fromField.Piece
	fromField.removePiece()
	toField.addPiece(piece)
}

func (b *Board) MovesForPiece(piece *Piece, ignorePins bool) *AvailableMoves {

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

	return &AvailableMoves{
		Moves: moves,
	}
}

func GetSquareSize() int {
	return squareSize
}

func GetBoardSize() int {
	return size
}

func (b *Board) MovesForColor(color string) *AvailableMoves {
	var moves []*Move

	for _, piece := range b.Pieces() {
		if piece.Color == color {
			moves = append(moves, b.MovesForPiece(piece, true).Moves...)
		}
	}

	return &AvailableMoves{
		Moves: moves,
	}
}

func (b *Board) isInCheck(color string) bool {
	var oppositeColor string
	if color == "white" {
		oppositeColor = "black"
	} else {
		oppositeColor = "white"
	}

	for _, move := range b.MovesForColor(oppositeColor).Moves {
		if move.To.Piece != nil && move.To.Piece.Type == "king" && move.To.Piece.Color == color {
			return true
		}
	}

	return false
}

func (b *Board) isInCheckAfterMove(move Move) bool {
	color := move.From.Piece.Color
	newBoard := b.deepCopy()
	newBoard.ExecuteMove(&move)
	return newBoard.isInCheck(color)
}

func (b *Board) IsCheckMate(color string) bool {
	// TODO Method is not working yet :D
	if b.isInCheck(color) {
		for _, piece := range b.Pieces() {
			if piece.Color == color {
				moves := b.MovesForPiece(piece, true).Moves
				if moves == nil || len(moves) > 0 {
					return false
				}
			}
		}
		return true
	}
	return false
}
