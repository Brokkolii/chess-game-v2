package main

import (
	"log"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/bot"
	"github.com/Brokkolii/chess-game-v2/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Type string
	Name string
}

type Game struct {
	State          *game.GameState
	DraggingPiece  *board.Piece
	AvailableMoves *board.AvailableMoves
	whitePlayer    *Player
	blackPlayer    *Player
}

func (g *Game) Update() error {
	if g.State.Turn == "white" {
		if g.whitePlayer.Type == "human" {
			handleUserInput(g)
		} else if g.whitePlayer.Type == "bot" {
			bot.PlayMove(g.State)
		}
	} else {
		if g.blackPlayer.Type == "human" {
			handleUserInput(g)
		} else if g.blackPlayer.Type == "bot" {
			bot.PlayMove(g.State)
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
		whitePlayer: &Player{
			Type: "bot",
			Name: "Brokkolii_1",
		},
		blackPlayer: &Player{
			Type: "bot",
			Name: "Brokkolii_Bot",
		},
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
