package chess

import (
	"fmt"
	"strings"
)

// Board is a tiny board representation, close to the Java version.
type Board struct {
	pieces      []*Piece
	playedMoves []Move
	captures    []*Piece
}

// NewBoardFromFEN creates a new board from a FEN-like string (only piece placement part).
func NewBoardFromFEN(fen string) (*Board, error) {
	b := &Board{
		pieces:      make([]*Piece, 120),
		playedMoves: []Move{},
		captures:    []*Piece{},
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
			if piece == nil || *piece == BLOCKER {
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

// CopyBoard creates a deep copy of the board.
func CopyBoard(src *Board) *Board {
	newPieces := make([]*Piece, len(src.pieces))
	copy(newPieces, src.pieces)
	newPlayedMoves := make([]Move, len(src.playedMoves))
	copy(newPlayedMoves, src.playedMoves)
	newCaptures := make([]*Piece, len(src.captures))
	copy(newCaptures, src.captures)
	return &Board{
		pieces:      newPieces,
		playedMoves: newPlayedMoves,
		captures:    newCaptures,
	}
}

func (b *Board) fillBlockers() {
	for i := 0; i < 20; i++ {
		b.pieces[i] = &BLOCKER
	}
	for i := 100; i < 120; i++ {
		b.pieces[i] = &BLOCKER
	}
	for i := 2; i < 10; i++ {
		startRank := i * 10
		b.pieces[startRank] = &BLOCKER
		b.pieces[startRank+9] = &BLOCKER
	}
}

// GetPiece returns the piece at a given square index.
func (b *Board) GetPiece(square int) *Piece {
	if square < 0 || square >= len(b.pieces) {
		return nil
	}
	return b.pieces[square]
}

// GetPieceByUCI returns the piece at a given UCI square (e.g. "e2").
func (b *Board) GetPieceByUCI(uci string) *Piece {
	return b.GetPiece(GetSquare(uci))
}

// SetPiece sets the piece at a given square index.
func (b *Board) SetPiece(square int, piece *Piece) {
	if square >= 0 && square < len(b.pieces) {
		b.pieces[square] = piece
	}
}

// GetMoves returns all moves for the given color.
func (b *Board) GetMoves(white bool) []Move {
	moves := []Move{}
	moveBuilder := moveConstructorFunc(func(from, to int) interface{} {
		return NewMove(from, to)
	})
	for square := 20; square < 100; square++ {
		piece := b.GetPiece(square)
		if piece != nil && *piece != BLOCKER && piece.IsWhite == white {
			mb := GetMoveBuilder(piece)
			if mb != nil {
				mb.Build(moves, b, square, moveBuilder)
			}
		}
	}
	return moves
}

// MakeMove applies a move to the board.
func (b *Board) MakeMove(move Move) error {
	if (move == Move{}) {
		return fmt.Errorf("move cannot be nil")
	}
	from := move.From()
	to := move.To()
	if from < 20 || from > 119 || to < 20 || to > 119 {
		return fmt.Errorf("illegal move")
	}
	movingPiece := b.pieces[from]
	capturedPiece := b.pieces[to]
	if movingPiece == nil || *movingPiece == BLOCKER || (capturedPiece != nil && !capturedPiece.CanBeCapturedBy(movingPiece.IsWhite)) {
		return fmt.Errorf("illegal move")
	}
	b.captures = append(b.captures, b.pieces[to])
	b.pieces[to] = b.pieces[from]
	b.pieces[from] = nil
	b.playedMoves = append(b.playedMoves, move)
	return nil
}

// UnmakeMove undoes the last move.
func (b *Board) UnmakeMove() error {
	if len(b.playedMoves) == 0 {
		return fmt.Errorf("no moves to unmake")
	}
	lastMove := b.playedMoves[len(b.playedMoves)-1]
	b.playedMoves = b.playedMoves[:len(b.playedMoves)-1]
	from := lastMove.From()
	to := lastMove.To()
	b.pieces[from] = b.pieces[to]
	b.pieces[to] = b.captures[len(b.captures)-1]
	b.captures = b.captures[:len(b.captures)-1]
	return nil
}

// Explorable interface implementation
func (b *Board) GetCapturable(square int) Capturable {
	return b.GetPiece(square)
}

func (b *Board) GetRank(square int) int {
	return getRank(square)
}

// getRank returns the rank (row) for a given square index.
func getRank(square int) int {
	return (square - 21) / 10
}

// GetSquare returns the board index for a UCI square (e.g. "e2").
func GetSquare(uci string) int {
	if len(uci) != 2 {
		panic("invalid UCI square: " + uci)
	}
	file := uci[0]
	rank := uci[1]
	if rank < '1' || rank > '8' || file < 'a' || file > 'h' {
		panic("invalid UCI square: " + uci)
	}
	return 21 + 10*int(rank-'1') + int(file-'a')
}

// GetUCI returns the UCI string for a board index.
func GetUCI(square int) string {
	square -= 21
	return fmt.Sprintf("%c%d", square%10+'a', square/10+1)
}

// String returns a string representation of the board.
func (b *Board) String() string {
	var sb strings.Builder
	for rank := 8; rank >= 1; rank-- {
		sb.WriteByte(byte('0' + rank))
		sb.WriteByte(' ')
		for file := 'a'; file <= 'h'; file++ {
			piece := b.GetPieceByUCI(fmt.Sprintf("%c%d", file, rank))
			if piece == nil {
				sb.WriteByte(' ')
			} else {
				sb.WriteRune(piece.Code)
			}
		}
		sb.WriteByte('\n')
	}
	sb.WriteString("  ")
	for file := 'a'; file < 'h'; file++ {
		sb.WriteByte(byte(file))
	}
	return sb.String()
}

// moveConstructorFunc is an adapter to allow the use of ordinary functions as MoveConstructor.
type moveConstructorFunc func(from, to int) interface{}

func (f moveConstructorFunc) Create(from, to int) interface{} {
	return f(from, to)
}
