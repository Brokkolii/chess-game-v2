package board

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Piece struct {
	Color string
	Type  string
	Image ebiten.Image
	X     int
	Y     int
	Row   int
	Col   int
}

func NewPiece(color string, pieceType string, row int, col int) *Piece {
	img, _, err := ebitenutil.NewImageFromFile("assets/" + color + "_" + pieceType + ".png")
	if err != nil {
		log.Fatal(err)
	}
	return &Piece{
		Color: color,
		Type:  pieceType,
		Image: *img,
		X:     col * squareSize,
		Y:     (size * squareSize) - ((row + 1) * squareSize),
		Row:   row,
		Col:   col,
	}
}

func DrawPieces(screen *ebiten.Image, board *Board) {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			piece := board[row][col]
			if piece != nil {
				drawPiece(board[row][col], screen)
			}
		}
	}
}

func drawPiece(piece *Piece, dst *ebiten.Image) {
	iw := float64(piece.Image.Bounds().Dx())
	ih := float64(piece.Image.Bounds().Dy())
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(squareSize/float64(iw), squareSize/float64(ih))
	op.GeoM.Translate(float64(piece.X), float64(piece.Y))
	dst.DrawImage(&piece.Image, &op)
}
