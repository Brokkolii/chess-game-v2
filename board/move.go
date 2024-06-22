package board

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Move struct {
	From *Field
	To   *Field
}

func NewMove(from *Field, to *Field) *Move {
	return &Move{
		From: from,
		To:   to,
	}
}

func (m *Move) Draw(dst *ebiten.Image) {
	radius := squareSize / 4
	offset := radius * 2
	fx, fy := m.To.getCoords()
	x := fx + offset
	y := fy + offset
	color := color.RGBA{111, 143, 114, 128}
	vector.DrawFilledCircle(dst, float32(x), float32(y), float32(radius), color, false)
}

func FieldInMoves(moves []*Move, field *Field) bool {
	for _, move := range moves {
		if move.To == field {
			return true
		}
	}
	return false
}

func (m *Move) IsPawnAdvance() bool {
	return m.From.Piece.Type == "pawn"
}

func (m *Move) IsCapture() bool {
	return m.To.Piece != nil
}

func (m *Move) AllowsEnPassant() bool {
	return m.From.Piece.Type == "pawn" && (m.From.Row-m.To.Row == -2 || m.From.Row-m.To.Row == 2)
}

func (m *Move) IsCastle() bool {
	return m.From.Piece.Type == "king" && (m.From.Col-m.To.Col == -3 || m.From.Col-m.To.Col == 3)
}

func (m *Move) IsFullMove() bool {
	// because white starts every a FullMove happens after blacks turn
	return m.From.Piece.Color == "black"
}
