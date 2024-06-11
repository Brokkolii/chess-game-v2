package board

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Fields [size][size]*Field
	Image  *ebiten.Image
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

func NewBoard() *Board {
	var board Board
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board.fromFEN(fen)
	return &board
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
	return b.FieldAt(row, col)
}

func (b *Board) FieldAt(row int, col int) *Field {
	return b.Fields[row-1][col-1]
}

func (b *Board) AddField(field *Field) {
	b.Fields[field.Row-1][field.Col-1] = field
}

func (b *Board) MovePiece(from *Field, to *Field) {
	piece := from.Piece
	from.removePiece()
	to.addPiece(piece)
}
