package board

import (
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

func GetSquareSize() int {
	return squareSize
}

func GetBoardSize() int {
	return size
}
