package board

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Piece struct {
	Color string
	Type  string
}

type Board [size][size]*Piece

const (
	size = 8
)

var (
	lightColor = color.RGBA{173, 189, 143, 0}
	darkColor  = color.RGBA{111, 143, 114, 0}
)

func DrawBoard(w int, h int, board *Board) *ebiten.Image {
	var boardImage = ebiten.NewImage(w, h)
	sw := w / size
	sh := h / size
	for row := 0; row < size; row++ {
		y := (size - 1 - row) * sh
		for col := 0; col < size; col++ {
			x := col * sw
			altColor := (row+col)%2 == 0
			piece := board[row][col]
			drawSquare(boardImage, x, y, sw, sh, altColor, piece)
		}
	}
	return boardImage
}

func drawSquare(screen *ebiten.Image, x int, y int, w int, h int, altColor bool, piece *Piece) {
	s := ebiten.NewImage(w, h)
	if !altColor {
		s.Fill(lightColor)
	} else {
		s.Fill(darkColor)
	}
	if piece != nil {
		drawPiece(piece, s)
	}
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(s, &op)
}

func drawPiece(piece *Piece, dst *ebiten.Image) {
	fontFace := basicfont.Face7x13
	x := 5
	y := dst.Bounds().Dy() / 2
	text.Draw(dst, piece.Color+" "+piece.Type, fontFace, x, y, color.Black)
}

func NewBoard() *Board {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	board := fenToBoard(fen)
	return board
}

func fenToBoard(fen string) *Board {
	var board Board
	row := 7
	col := 0
	for _, char := range fen {
		switch char {
		case 'r':
			board[row][col] = &Piece{Color: "black", Type: "rook"}
			col++
		case 'n':
			board[row][col] = &Piece{Color: "black", Type: "knight"}
			col++
		case 'b':
			board[row][col] = &Piece{Color: "black", Type: "bishop"}
			col++
		case 'q':
			board[row][col] = &Piece{Color: "black", Type: "queen"}
			col++
		case 'k':
			board[row][col] = &Piece{Color: "black", Type: "king"}
			col++
		case 'p':
			board[row][col] = &Piece{Color: "black", Type: "pawn"}
			col++
		case 'R':
			board[row][col] = &Piece{Color: "white", Type: "rook"}
			col++
		case 'N':
			board[row][col] = &Piece{Color: "white", Type: "knight"}
			col++
		case 'B':
			board[row][col] = &Piece{Color: "white", Type: "bishop"}
			col++
		case 'Q':
			board[row][col] = &Piece{Color: "white", Type: "queen"}
			col++
		case 'K':
			board[row][col] = &Piece{Color: "white", Type: "king"}
			col++
		case 'P':
			board[row][col] = &Piece{Color: "white", Type: "pawn"}
			col++
		case '/':
			row--
			col = 0
		case ' ':
			return &board
		default:
			if char >= '1' && char <= '8' {
				col += int(char - '0')
			}
		}
	}
	return &board
}
