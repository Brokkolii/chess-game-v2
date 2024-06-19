package main

import (
	"fmt"
	"log"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/bot"
	"github.com/Brokkolii/chess-game-v2/match"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	State          *match.State
	DraggingPiece  *board.Piece
	AvailableMoves *board.AvailableMoves
}

func (g *Game) Update() error {
	if g.State.MatchIsRunning() {
		turnType := g.State.Turn.PlayerType
		var move *board.Move
		if turnType == "human" {
			move = GetUserMove(g)
		} else if turnType == "bot" {
			move = bot.GetMove(g.State)
		}

		if move != nil {
			g.State.PlayMove(move)
		}
	} else {
		// TODO Temp solution until there is a menu
		fmt.Println("Starting new Game")
		g.State.StartMatch()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State.MatchIsRunning() {
		g.State.Board.Draw(screen)
		g.State.Board.DrawPieces(screen)
		if g.AvailableMoves != nil {
			g.State.Board.DrawMoves(screen, g.AvailableMoves.Moves)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func NewGame() *Game {
	return &Game{
		State: match.NewState(),
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
