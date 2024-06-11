package main

import (
	"log"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	State         *game.GameState
	DraggingPiece *board.Piece
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.DraggingPiece == nil {
			clickedPiece := g.State.Board.PieceAtCoords(mx, my)
			if clickedPiece != nil {
				clickedPiece.IsDragged = true
				g.DraggingPiece = clickedPiece
			}
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.DraggingPiece != nil {
			field := g.State.Board.FieldAtCoords(mx, my)
			// TODO: check if valid move
			g.State.Board.MovePiece(g.DraggingPiece.Field, field)
			g.DraggingPiece.IsDragged = false
			g.DraggingPiece = nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.State.Board.Draw(screen)
	g.State.Board.DrawPieces(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func NewGame() *Game {
	return &Game{
		State: game.NewGameState(),
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
