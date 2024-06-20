package board

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	lightColor = color.RGBA{173, 189, 143, 0}
	darkColor  = color.RGBA{111, 143, 114, 0}
)

type Field struct {
	Row   int // 1-8
	Col   int // 1-8
	Piece *Piece
}

func (f *Field) getLabel() string {
	return fmt.Sprint(f.Row) + fmt.Sprint(string(rune('A'+f.Col-1)))
}

func (f *Field) getCoords() (x int, y int) {
	x = (f.Col - 1) * squareSize
	y = (size - f.Row) * squareSize
	return x, y
}

func (f *Field) GetColor() *color.RGBA {
	isLight := ((f.Row+f.Col)%2 == 0)
	if isLight {
		return &lightColor
	} else {
		return &darkColor
	}
}

func (f *Field) draw(dst *ebiten.Image) {
	s := ebiten.NewImage(squareSize, squareSize)
	color := f.GetColor()
	s.Fill(color)
	x, y := f.getCoords()
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	dst.DrawImage(s, &op)
}

func (f *Field) addPiece(piece *Piece) {
	f.Piece = piece
	piece.Field = f
}

func (f *Field) removePiece() *Piece {
	piece := f.Piece
	f.Piece = nil
	piece.Field = nil
	return piece
}

func NewField(row int, col int) *Field {
	return &Field{
		Row: row,
		Col: col,
	}
}

func NewFieldWithPiece(row int, col int, piece *Piece) *Field {
	field := NewField(row, col)
	field.addPiece(piece)
	return field
}
