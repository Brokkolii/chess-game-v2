package main

import (
	"log"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	State          *game.GameState
	DraggingPiece  *board.Piece
	AvailableMoves *board.AvailableMoves
}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.DraggingPiece == nil {
			clickedPiece := g.State.Board.PieceAtCoords(mx, my)
			if clickedPiece != nil {
				clickedPiece.IsDragged = true
				g.DraggingPiece = clickedPiece
				g.AvailableMoves = g.State.Board.MovesForPiece(clickedPiece)
			}
		}
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if g.DraggingPiece != nil {
			field := g.State.Board.FieldAtCoords(mx, my)
			if field != nil {
				// TODO: check if valid move
				isAvailable := g.AvailableMoves.FieldIsAvailable(field)
				if isAvailable {
					move := board.NewMove(g.DraggingPiece.Field, field)
					g.State.Board.ExecuteMove(move)
				}
			}
			g.DraggingPiece.IsDragged = false
			g.DraggingPiece = nil
			g.AvailableMoves = nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.State.Board.Draw(screen)
	g.State.Board.DrawPieces(screen)
	if g.AvailableMoves != nil {
		g.State.Board.DrawMoves(screen, g.AvailableMoves.Moves)
	}
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
