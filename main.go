package main

import (
	"log"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	State      *game.GameState
	BoardImage *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(board.DrawBoard(640, 640, g.State.Board), nil)
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
