package main

import (
	"log"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	State      *game.GameState
	BoardImage *ebiten.Image
	Dragging   *board.Piece
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.Dragging == nil {
			clickedPiece := g.State.Board.PieceAt(mx, my)
			if clickedPiece != nil {
				g.Dragging = clickedPiece
			}
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.Dragging != nil {
			row, col := g.State.Board.SquareAt(mx, my)
			g.State.Board.MovePiece(g.Dragging.Row, g.Dragging.Col, row, col)
			g.Dragging = nil
		}
	}
	if g.Dragging != nil {
		g.Dragging.X = mx - board.GetSquareSize()/2
		g.Dragging.Y = my - board.GetSquareSize()/2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.BoardImage, nil)
	board.DrawPieces(screen, g.State.Board)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func NewGame() *Game {
	return &Game{
		State:      game.NewGameState(),
		BoardImage: board.DrawBoard(640, 640),
	}
}

func main() {
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Chess Game V2")
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
