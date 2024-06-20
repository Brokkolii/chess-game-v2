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
