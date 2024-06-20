package match

import (
	"fmt"

	"github.com/Brokkolii/chess-game-v2/board"
	"github.com/Brokkolii/chess-game-v2/util"
)

type Player struct {
	PlayerType string
	Name       string
	Color      string
}

type State struct {
	Board        *board.Board
	Turn         *Player
	player1      *Player
	player2      *Player
	matchStarted bool
	matchOver    bool
	winner       *Player
}

func NewState() *State {
	return &State{
		Board:        board.NewBoard(),
		Turn:         nil,
		player1:      nil,
		player2:      nil,
		matchStarted: false,
		matchOver:    false,
		winner:       nil,
	}
}

func (s *State) StartMatch() {
	player1 := &Player{
		PlayerType: "human",
		Name:       "Brokkolii_1",
		Color:      "white",
	}
	s.player1 = player1

	player2 := &Player{
		PlayerType: "bot",
		Name:       "Brokkolii_Bot",
		Color:      "black",
	}
	s.player2 = player2

	s.matchOver = false
	s.winner = nil
	s.Turn = s.getPlayerWithColor("white")
	s.matchStarted = true
	s.Board = board.NewBoard()
}

func (s *State) endMatch(winner *Player) {
	fmt.Println("Match is over!", winner.Color, "won!")
	s.winner = winner
	s.matchOver = true
	s.Turn = nil
}

func (s *State) nextTurn() {
	if s.Turn == s.player1 {
		s.Turn = s.player2
	} else {
		s.Turn = s.player1
	}
}

func (s *State) PlayMove(move *board.Move) {
	if move.From.Piece.Color == s.Turn.Color {
		s.Board.ExecuteMove(move)

		enemyColor := util.InvertColor(s.Turn.Color)
		if s.Board.IsCheckMate(enemyColor) {
			s.endMatch(s.Turn)
		}

		s.nextTurn()
	}
}

func (s *State) getPlayerWithColor(Color string) *Player {
	if s.player1 != nil && s.player1.Color == Color {
		return s.player1
	} else if s.player2 != nil && s.player2.Color == Color {
		return s.player2
	} else {
		return nil
	}
}

func (s *State) MatchIsRunning() bool {
	return s.matchStarted && !s.matchOver
}
