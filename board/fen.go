package board

import (
	"fmt"
	"strconv"
	"strings"
)

func (b *Board) ToFEN() string {
	fen := ""

	// setup board
	epmtyFieldsCounter := 0
	for row := 8; row >= 1; row-- {
		for col := 1; col <= 8; col++ {
			piece := b.FieldAt(row, col).Piece
			if piece != nil {
				if epmtyFieldsCounter > 0 {
					fen = fen + fmt.Sprint(epmtyFieldsCounter)
				}
				epmtyFieldsCounter = 0
				// ad letter
				if piece.Color == "white" {
					switch piece.Type {
					case "rook":
						fen = fen + "R"
					case "bishop":
						fen = fen + "B"
					case "knight":
						fen = fen + "N"
					case "queen":
						fen = fen + "Q"
					case "king":
						fen = fen + "K"
					case "pawn":
						fen = fen + "P"
					}
				} else {
					switch piece.Type {
					case "rook":
						fen = fen + "r"
					case "bishop":
						fen = fen + "b"
					case "knight":
						fen = fen + "n"
					case "queen":
						fen = fen + "q"
					case "king":
						fen = fen + "k"
					case "pawn":
						fen = fen + "p"
					}
				}
			} else {
				epmtyFieldsCounter++
			}
		}
		if epmtyFieldsCounter > 0 {
			fen = fen + fmt.Sprint(epmtyFieldsCounter)
		}
		epmtyFieldsCounter = 0
		fen = fen + "/"
	}
	fen += " "

	// turn
	if b.Turn == "white" {
		fen += "w"
	} else {
		fen += "b"
	}
	fen += " "

	// castling
	castling := ""
	if b.BlackKingsideCastle {
		castling += "K"
	}
	if b.BlackQueensideCastle {
		castling += "Q"
	}
	if b.WhiteKingsideCastle {
		castling += "k"
	}
	if b.WhiteQueensideCastle {
		castling += "q"
	}
	if castling == "" {
		castling = "-"
	}
	fen += castling + " "

	// enpassante
	// TODO implement enpassante
	fen += "- "

	// half move clock
	fen += strconv.Itoa(b.HalfMoveClock) + " "

	// full move number
	fen += strconv.Itoa(b.FullMoveNumber) + " "

	return fen
}

func (b *Board) FromFEN(fen string) error {
	fenParts := strings.Split(fen, " ")

	// setup board
	row := 8
	col := 1
	for _, char := range fenParts[0] {
		switch char {
		case 'r':
			piece := NewPiece("black", "rook", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'n':
			piece := NewPiece("black", "knight", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'b':
			piece := NewPiece("black", "bishop", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'q':
			piece := NewPiece("black", "queen", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'k':
			piece := NewPiece("black", "king", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'p':
			piece := NewPiece("black", "pawn", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'R':
			piece := NewPiece("white", "rook", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'N':
			piece := NewPiece("white", "knight", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'B':
			piece := NewPiece("white", "bishop", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'Q':
			piece := NewPiece("white", "queen", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'K':
			piece := NewPiece("white", "king", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case 'P':
			piece := NewPiece("white", "pawn", row, col)
			field := NewFieldWithPiece(row, col, piece)
			b.AddField(field)
			col++
		case '/':
			row--
			col = 1
		default:
			if char >= '1' && char <= '8' {
				for i := 1; i <= int(char-'0'); i++ {
					field := NewField(row, col)
					b.AddField(field)
					col++
				}
			}
		}
	}

	// turn
	if fenParts[1] == "w" {
		b.Turn = "white"
	} else {
		b.Turn = "black"
	}

	// castling
	b.BlackKingsideCastle = strings.Contains(fenParts[2], "K")
	b.BlackQueensideCastle = strings.Contains(fenParts[2], "Q")
	b.WhiteKingsideCastle = strings.Contains(fenParts[2], "k")
	b.WhiteQueensideCastle = strings.Contains(fenParts[2], "q")

	//en passante
	// TODO implement en passante
	b.EnPassant = nil

	// half move clock
	hmc, err := strconv.Atoi(fenParts[4])
	if err != nil {
		return err
	}
	b.HalfMoveClock = hmc

	// half move clock
	fmn, err := strconv.Atoi(fenParts[5])
	if err != nil {
		return err
	}
	b.FullMoveNumber = fmn

	return nil
}
