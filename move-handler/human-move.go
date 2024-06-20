package movehandler

import (
	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/match"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (mh *MoveHandler) CheckForHumanMove(s *match.State) {
	if s.Turn.PlayerType != "human" {
		return
	}
	// get available moves, pieceMoves, consideringMoves
	// and choosenMove => then stop thinking
	mc := mh.moveCollection

	// LEFT MOUSE BUTTON PRESSED
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if mh.clickedPiece == nil {
			mx, my := ebiten.CursorPosition()
			clickedPiece := s.Board.PieceAtCoords(mx, my)
			if clickedPiece != nil && clickedPiece.Color == s.Turn.Color {
				clickedPiece.IsDragged = true
				mh.clickedPiece = clickedPiece

				go func() {
					moves := s.Board.MovesForPiece(clickedPiece, false)
					if mh.clickedPiece == clickedPiece {
						mc.pieceMoves = moves
					}
				}()
			}
		}
	}

	// LEFT MOUSE BUTTON UP
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		if mh.clickedPiece != nil {
			mx, my := ebiten.CursorPosition()
			field := s.Board.FieldAtCoords(mx, my)
			if field != nil {
				if board.FieldInMoves(mc.pieceMoves, field) {
					move := board.NewMove(mh.clickedPiece.Field, field)
					mh.moveCollection.availableMoves = nil
					mh.moveCollection.consideringMove = nil
					mh.moveCollection.choosenMove = move
					mh.thinking = false
				}
			}
			mh.clickedPiece.IsDragged = false
			mh.clickedPiece = nil
			mh.moveCollection.pieceMoves = nil
		}
	}
}
