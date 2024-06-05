package board

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board [size][size]*Piece

const (
	size       = 8
	squareSize = 80
)

var (
	lightColor = color.RGBA{173, 189, 143, 0}
	darkColor  = color.RGBA{111, 143, 114, 0}
)

func DrawBoard(w int, h int) *ebiten.Image {
	var boardImage = ebiten.NewImage(w, h)
	for row := 0; row < size; row++ {
		y := (size - 1 - row) * squareSize
		for col := 0; col < size; col++ {
			x := col * squareSize
			altColor := (row+col)%2 == 0
			drawSquare(boardImage, x, y, altColor)
		}
	}
	return boardImage
}

func drawSquare(screen *ebiten.Image, x int, y int, altColor bool) {
	s := ebiten.NewImage(squareSize, squareSize)
	if !altColor {
		s.Fill(lightColor)
	} else {
		s.Fill(darkColor)
	}
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(s, &op)
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
			board[row][col] = NewPiece("black", "rook", row, col)
			col++
		case 'n':
			board[row][col] = NewPiece("black", "knight", row, col)
			col++
		case 'b':
			board[row][col] = NewPiece("black", "bishop", row, col)
			col++
		case 'q':
			board[row][col] = NewPiece("black", "queen", row, col)
			col++
		case 'k':
			board[row][col] = NewPiece("black", "king", row, col)
			col++
		case 'p':
			board[row][col] = NewPiece("black", "pawn", row, col)
			col++
		case 'R':
			board[row][col] = NewPiece("white", "rook", row, col)
			col++
		case 'N':
			board[row][col] = NewPiece("white", "knight", row, col)
			col++
		case 'B':
			board[row][col] = NewPiece("white", "bishop", row, col)
			col++
		case 'Q':
			board[row][col] = NewPiece("white", "queen", row, col)
			col++
		case 'K':
			board[row][col] = NewPiece("white", "king", row, col)
			col++
		case 'P':
			board[row][col] = NewPiece("white", "pawn", row, col)
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

func (b *Board) Pieces() []*Piece {
	var pieces []*Piece
	for _, row := range b {
		for _, piece := range row {
			if piece != nil {
				pieces = append(pieces, piece)
			}
		}
	}
	return pieces
}

func (b *Board) PieceAt(x int, y int) *Piece {
	for _, piece := range b.Pieces() {
		if x-piece.X <= squareSize && y-piece.Y <= squareSize && x >= piece.X && y >= piece.Y {
			return piece
		}
	}
	return nil
}

func (b *Board) SquareAt(x int, y int) (row int, col int) {
	return size - 1 - (y / squareSize), x / squareSize
}

func (b *Board) MovePiece(fromRow int, fromCol int, toRow int, toCol int) {
	b[toRow][toCol] = b[fromRow][fromCol]
	if toRow != fromRow || toCol != fromCol {
		b[fromRow][fromCol] = nil
	}
	piece := b[toRow][toCol]
	piece.Row = toRow
	piece.Col = toCol
	piece.X = piece.Col * squareSize
	piece.Y = (size * squareSize) - ((piece.Row + 1) * squareSize)
}

func GetSquareSize() int {
	return squareSize
}
