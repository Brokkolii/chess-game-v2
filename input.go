package main

import (
	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// TODO: Horribe sideeffects with dragging piece
// 	=> 	need input handler struct that has the dragging
//		piece and maybe the marked moves as property

func GetUserMove(game *Game) *board.Move {
	mx, my := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		handleLeftMouseDown(game, mx, my)
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		move := handleLeftMouseUp(game, mx, my)
		if move != nil {
			return move
		}
	}
	return nil
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

func handleLeftMouseUp(game *Game, mx int, my int) *board.Move {
	if game.DraggingPiece != nil {
		field := game.State.Board.FieldAtCoords(mx, my)
		if field != nil {
			isAvailable := game.AvailableMoves.FieldIsAvailable(field)
			if isAvailable {
				move := board.NewMove(game.DraggingPiece.Field, field)
				game.DraggingPiece.IsDragged = false
				game.DraggingPiece = nil
				game.AvailableMoves = nil
				return move
			}
		}
		game.DraggingPiece.IsDragged = false
		game.DraggingPiece = nil
		game.AvailableMoves = nil
	}
	return nil
}
