package movehandler

import (
	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/match"
)

type MoveCollection struct {
	availableMoves  []*board.Move
	pieceMoves      []*board.Move
	consideringMove *board.Move
	choosenMove     *board.Move
}

type MoveHandler struct {
	thinking       bool
	clickedPiece   *board.Piece
	moveCollection *MoveCollection
}

func NewMoveCollection() *MoveCollection {
	return &MoveCollection{
		availableMoves:  nil,
		pieceMoves:      nil,
		consideringMove: nil,
		choosenMove:     nil,
	}
}

func NewMoveHandler() *MoveHandler {
	return &MoveHandler{
		thinking:       false,
		moveCollection: NewMoveCollection(),
	}
}

func (mh *MoveHandler) Reset() {
	mh.thinking = false
	mh.clickedPiece = nil
	mh.moveCollection.availableMoves = nil
	mh.moveCollection.pieceMoves = nil
	mh.moveCollection.consideringMove = nil
	mh.moveCollection.choosenMove = nil
}

func (mh *MoveHandler) RequestMove(state *match.State) {
	mh.Reset()
	mh.thinking = true

	go func() {
		moves := state.Board.MovesForColor(state.Turn.Color, false)
		if mh.thinking {
			mh.moveCollection.availableMoves = moves
		}
	}()

	if state.Turn.PlayerType == "bot" {
		c := make(chan *MoveCollection)
		go mh.waitForBotMove(state, c)
		//Receive messages from the channel
		for mc := range c {
			mh.moveCollection = mc
			mh.thinking = mh.moveCollection == nil
		}
	}

}

func (mh *MoveHandler) IsThinking() bool {
	return mh.thinking
}

func (mh *MoveHandler) GetAvailableMoves() []*board.Move {
	return mh.moveCollection.availableMoves
}

func (mh *MoveHandler) GetPieceMoves() []*board.Move {
	return mh.moveCollection.pieceMoves
}

func (mh *MoveHandler) GetConsideringMove() *board.Move {
	return mh.moveCollection.consideringMove
}

func (mh *MoveHandler) GetChoosenMove() *board.Move {
	return mh.moveCollection.choosenMove
}
