package main

import (
	"fmt"
	"log"

	"github.com/Brokkolii/chess-game-v2/match"
	movehandler "github.com/Brokkolii/chess-game-v2/move-handler"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	State       *match.State
	MoveHandler *movehandler.MoveHandler
}

func (g *Game) Update() error {
	if g.State.MatchIsRunning() {
		if !g.MoveHandler.IsThinking() {
			move := g.MoveHandler.GetChoosenMove()
			if move != nil {
				g.State.PlayMove(move)
				g.MoveHandler.Reset()
			} else {
				g.MoveHandler.RequestMove(g.State)
			}

		} else {
			g.MoveHandler.CheckForHumanMove(g.State)
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

		// TODO draw movecollection of movehandler
		// TODO rethink drawMoves func => infavor of marking Fields in ways
		pieceMoves := g.MoveHandler.GetPieceMoves()
		if pieceMoves != nil {
			g.State.Board.DrawMoves(screen, pieceMoves)
		}

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func NewGame() *Game {
	return &Game{
		State:       match.NewState(),
		MoveHandler: movehandler.NewMoveHandler(),
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
