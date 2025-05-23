package chess

import (
	"fmt"
	"strings"
)

// Board represents a chess board.
type Board struct {
	pieces      []*Piece
	playedMoves []Move
	captures    []*Piece
}

// NewBoard creates a new board from a FEN string.
func NewBoard(fen string) (*Board, error) {
	b := &Board{
		pieces:      make([]*Piece, 120),
		playedMoves: make([]Move, 0),
		captures:    make([]*Piece, 0),
	}
	b.fillBlockers()

	rank := 7
	file := 0
	for _, c := range fen {
		switch {
		case c >= '1' && c <= '8':
			count := int(c - '0')
			if count > 8-file {
				return nil, fmt.Errorf("invalid FEN: too many pieces on rank %d", rank)
			}
			file += count
		case c == '/':
			if file != 8 {
				return nil, fmt.Errorf("invalid FEN: missing files on rank %d", rank)
			}
			if rank == 0 {
				return nil, fmt.Errorf("invalid FEN: too many ranks")
			}
			file = 0
			rank--
		default:
			piece := FromCode(c)
			if piece == nil || piece == &BLOCKER {
				return nil, fmt.Errorf("invalid FEN: unknown piece %c", c)
			}
			b.pieces[21+rank*10+file] = piece
			file++
		}
	}

	if file != 8 {
		return nil, fmt.Errorf("invalid FEN: missing files on rank %d", rank)
	}
	if rank != 0 {
		return nil, fmt.Errorf("invalid FEN: missing ranks")
	}

	return b, nil
}

// NewBoardCopy creates a copy of an existing board.
func NewBoardCopy(src *Board) *Board {
	dst := &Board{
		pieces:      make([]*Piece, len(src.pieces)),
		playedMoves: make([]Move, len(src.playedMoves)),
		captures:    make([]*Piece, len(src.captures)),
	}
	copy(dst.pieces, src.pieces)
	copy(dst.playedMoves, src.playedMoves)
	copy(dst.captures, src.captures)
	return dst
}

// fillBlockers fills the board with blocker pieces around the edges.
func (b *Board) fillBlockers() {
	// Fill top and bottom borders
	for i := 0; i < 20; i++ {
		b.pieces[i] = &BLOCKER
		b.pieces[100+i] = &BLOCKER
	}

	// Fill left and right borders
	for i := 2; i < 10; i++ {
		b.pieces[i*10] = &BLOCKER
		b.pieces[i*10+9] = &BLOCKER
	}
}

// GetSquare converts a UCI square string to a board index.
func GetSquare(uciSquare string) int {
	if len(uciSquare) != 2 {
		panic(fmt.Sprintf("invalid UCI square: %s", uciSquare))
	}
	file := int(uciSquare[0] - 'a')
	rank := int(uciSquare[1] - '1')
	if rank < 0 || rank > 7 || file < 0 || file > 7 {
		panic(fmt.Sprintf("invalid UCI square: %s", uciSquare))
	}
	return 21 + 10*rank + file
}

// GetRank returns the rank (0-7) of a square.
func GetRank(square int) int {
	return (square - 21) / 10
}

// GetUCI returns the UCI notation for a square.
func GetUCI(square int) string {
	square -= 21
	file := byte(square%10) + 'a'
	rank := byte(square/10) + '1'
	return string([]byte{file, rank})
}

// GetPiece returns the piece at the given UCI square.
func (b *Board) GetPiece(uciSquare string) *Piece {
	return b.pieces[GetSquare(uciSquare)]
}

// GetPieceByIndex returns the piece at the given board index.
func (b *Board) GetPieceByIndex(square int) *Piece {
	return b.pieces[square]
}

// GetMoves returns all legal moves for the given color.
func (b *Board) GetMoves(white bool) []Move {
	moves := make([]Move, 0, 40) // Pre-allocate reasonable capacity

	for square := 20; square < 100; square++ {
		piece := b.pieces[square]
		if piece != nil && piece != &BLOCKER && piece.IsWhite == white {
			generator := GetMoveGenerator(piece)
			if generator == nil {
				fmt.Printf("No generator found for piece %c at %s\n", piece.Code, GetUCI(square))
			} else {
				beforeLen := len(moves)
				generator.Build(&moves, b, square)
				if len(moves) > beforeLen {
					fmt.Printf("Generated %d moves for %c at %s: %v\n", 
						len(moves) - beforeLen, piece.Code, GetUCI(square), 
						moves[beforeLen:])
				} else {
					fmt.Printf("No moves generated for %c at %s\n", piece.Code, GetUCI(square))
				}
			}
		}
	}

	return moves
}

// MakeMove makes a move on the board.
func (b *Board) MakeMove(move Move) error {
	if move == (Move{}) {
		return fmt.Errorf("move cannot be zero value")
	}

	from := move.From()
	to := move.To()

	if from < 20 || from > 119 || to < 20 || to > 119 {
		return fmt.Errorf("illegal move: out of bounds")
	}

	movingPiece := b.pieces[from]
	if movingPiece == nil || movingPiece == &BLOCKER {
		return fmt.Errorf("illegal move: no piece at source square")
	}

	capturedPiece := b.pieces[to]
	if capturedPiece != nil && !capturedPiece.CanBeCapturedBy(movingPiece.IsWhite) {
		return fmt.Errorf("illegal move: cannot capture own piece")
	}

	// Store the captured piece (if any)
	b.captures = append(b.captures, capturedPiece)


	// Move the piece
	b.pieces[to] = movingPiece
	b.pieces[from] = nil

	// Record the move
	b.playedMoves = append(b.playedMoves, move)
	return nil
}

// UnmakeMove undoes the last move made on the board.
func (b *Board) UnmakeMove() error {
	if len(b.playedMoves) == 0 {
		return fmt.Errorf("no moves to unmake")
	}

	// Get the last move and remove it from history
	lastMove := b.playedMoves[len(b.playedMoves)-1]
	b.playedMoves = b.playedMoves[:len(b.playedMoves)-1]

	from := lastMove.From()
	to := lastMove.To()

	// Restore the moved piece
	b.pieces[from] = b.pieces[to]

	// Restore the captured piece (if any)
	if len(b.captures) > 0 {
		b.pieces[to] = b.captures[len(b.captures)-1]
		b.captures = b.captures[:len(b.captures)-1]
	} else {
		b.pieces[to] = nil
	}

	return nil
}

// String returns a string representation of the board.
func (b *Board) String() string {
	var sb strings.Builder

	for rank := 7; rank >= 0; rank-- {
		sb.WriteByte(byte('1' + rank))
		sb.WriteByte(' ')
		for file := 0; file < 8; file++ {
			square := 21 + rank*10 + file
			piece := b.pieces[square]
			if piece == nil {
				sb.WriteByte(' ')
			} else {
				sb.WriteRune(piece.Code)
			}
		}
		sb.WriteByte('\n')
	}

	sb.WriteString("  ")
	for file := 0; file < 8; file++ {
		sb.WriteByte(byte('a' + file))
	}

	return sb.String()
}
