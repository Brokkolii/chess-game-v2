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
	player1      *Player
	player2      *Player
	matchStarted bool
	matchOver    bool
	winner       *Player
}

func NewState() *State {
	return &State{
		Board:        nil,
		player1:      nil,
		player2:      nil,
		matchStarted: false,
		matchOver:    false,
		winner:       nil,
	}
}

func (s *State) StartMatch() {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
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
	s.matchStarted = true
	s.Board = board.NewBoardFromFEN(fen)
}

func (s *State) endMatch(winner *Player) {
	fmt.Println("Match is over!", winner.Color, "won!")
	s.winner = winner
	s.matchOver = true
}

func (s *State) PlayMove(move *board.Move) {
	if move.From.Piece.Color == s.Board.Turn {
		if move.IsCapture() || move.IsPawnAdvance() {
			s.Board.HalfMoveClock = 0
		} else {
			s.Board.HalfMoveClock++
		}

		if move.IsFullMove() {
			s.Board.FullMoveNumber++
		}

		if move.IsCastle() {
			if move.From.Piece.Color == "black" {
				if move.To.Col == 8 {
					s.Board.BlackKingsideCastle = false
				} else {
					s.Board.BlackQueensideCastle = false
				}
			} else {
				if move.To.Col == 8 {
					s.Board.WhiteKingsideCastle = false
				} else {
					s.Board.WhiteQueensideCastle = false
				}
			}
		}

		if move.AllowsEnPassant() {
			// TODO: Implement EnPassant
		}

		s.Board.ExecuteMove(move)

		enemyColor := util.InvertColor(s.Board.Turn)
		if s.Board.IsCheckMate(enemyColor) {
			s.endMatch(s.GetPlayerWithColor(s.Board.Turn))
		}

		fmt.Println("new evaluation for", s.Board.Turn, s.Board.Evaluate())
		s.Board.NextTurn()

		fmt.Println("FEN for Board:", s.Board.ToFEN())
	}
}

func (s *State) GetPlayerWithColor(Color string) *Player {
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
