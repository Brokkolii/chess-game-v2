package board

import (
	"github.com/Brokkolii/chess-game-v2/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type Piece struct {
	Color     string
	Type      string
	Image     ebiten.Image
	Field     *Field
	IsDragged bool
}

func NewPiece(pieceColor string, pieceType string, row int, col int) *Piece {
	img := assets.LoadImageFromAssets(pieceColor + "_" + pieceType + ".png")
	return &Piece{
		Color:     pieceColor,
		Type:      pieceType,
		Image:     *img,
		IsDragged: false,
	}
}

func (p *Piece) Draw(dst *ebiten.Image) {
	iw := float64(p.Image.Bounds().Dx())
	ih := float64(p.Image.Bounds().Dy())
	x, y := p.getCoords()
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(squareSize/float64(iw), squareSize/float64(ih))
	op.GeoM.Translate(float64(x), float64(y))
	dst.DrawImage(&p.Image, &op)
}

func (p *Piece) getCoords() (x int, y int) {
	if p.IsDragged {
		x, y = ebiten.CursorPosition()
		x = x - (squareSize / 2)
		y = y - (squareSize / 2)
		return x, y
	} else {
		return p.Field.getCoords()
	}
}
