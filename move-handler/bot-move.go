package movehandler

import (
	"github.com/Brokkolii/chess-game-v2/bot"
	"github.com/Brokkolii/chess-game-v2/match"
)

func (mh *MoveHandler) waitForBotMove(state *match.State, c chan *MoveCollection) {
	mc := NewMoveCollection()
	move := bot.GetMove(state)
	mc.choosenMove = move
	c <- mc
	close(c)
}
