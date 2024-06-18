package main

import (
	"fmt"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func handleUserInput(game *Game) {
	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		handleLeftMouseDown(game, mx, my)
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		handleLeftMouseUp(game, mx, my)
	}
}

func handleLeftMouseDown(game *Game, mx int, my int) {
	if game.DraggingPiece == nil {
		clickedPiece := game.State.Board.PieceAtCoords(mx, my)
		if clickedPiece != nil {
			clickedPiece.IsDragged = true
			game.DraggingPiece = clickedPiece
			game.AvailableMoves = game.State.Board.MovesForPiece(clickedPiece, false)
		}
	}
}

func handleLeftMouseUp(game *Game, mx int, my int) {
	if game.DraggingPiece != nil {
		field := game.State.Board.FieldAtCoords(mx, my)
		if field != nil {
			isAvailable := game.AvailableMoves.FieldIsAvailable(field)
			if isAvailable {
				move := board.NewMove(game.DraggingPiece.Field, field)

				game.State.Board.ExecuteMove(move)
				game.State.NextTurn()

				fmt.Println("checking for mate")
				// CHECK FOR CHECKMATE
				if game.State.Board.IsCheckMate("white") {
					fmt.Println("BLACK WON")
				} else if game.State.Board.IsCheckMate("black") {
					fmt.Println("WHITE WON")
				}
			}
		}
		game.DraggingPiece.IsDragged = false
		game.DraggingPiece = nil
		game.AvailableMoves = nil
	}
}
