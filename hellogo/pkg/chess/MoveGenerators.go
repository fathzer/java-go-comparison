package chess

import (
	"fmt"
)

// ROW_WIDTH is the width of a row on the board.
const ROW_WIDTH = 10

// Direction represents a movement direction with its delta.
type Direction struct {
	delta int
}

func (d Direction) Delta() int {
	return d.delta
}

var (
	NORTH      = Direction{ROW_WIDTH}
	SOUTH      = Direction{-ROW_WIDTH}
	EAST       = Direction{+1}
	WEST       = Direction{-1}
	NORTH_EAST = Direction{ROW_WIDTH + 1}
	NORTH_WEST = Direction{ROW_WIDTH - 1}
	SOUTH_EAST = Direction{-ROW_WIDTH + 1}
	SOUTH_WEST = Direction{-ROW_WIDTH - 1}
)

// MoveGenerators holds move builder singletons.
// TODO This should probably be in a separate file with explicit reference to Move as ChatGPT has a lot of problem with it
type MoveGenerators struct{}

var (
	whiteKingMoveBuilder   = NewKingMoveBuilder[interface{}](true)
	whiteQueenMoveBuilder  = NewSliderMoveBuilder[interface{}]([]Direction{NORTH, SOUTH, EAST, WEST, NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST}, true)
	whiteRookMoveBuilder   = NewSliderMoveBuilder[interface{}]([]Direction{NORTH, SOUTH, EAST, WEST}, true)
	whiteBishopMoveBuilder = NewSliderMoveBuilder[interface{}]([]Direction{NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST}, true)
	whiteKnightMoveBuilder = NewKnightMoveBuilder[interface{}](true)
	whitePawnMoveBuilder   = NewPawnMoveBuilder[interface{}](true)

	blackKingMoveBuilder   = NewKingMoveBuilder[interface{}](false)
	blackQueenMoveBuilder  = NewSliderMoveBuilder[interface{}]([]Direction{NORTH, SOUTH, EAST, WEST, NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST}, false)
	blackRookMoveBuilder   = NewSliderMoveBuilder[interface{}]([]Direction{NORTH, SOUTH, EAST, WEST}, false)
	blackBishopMoveBuilder = NewSliderMoveBuilder[interface{}]([]Direction{NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST}, false)
	blackKnightMoveBuilder = NewKnightMoveBuilder[interface{}](false)
	blackPawnMoveBuilder   = NewPawnMoveBuilder[interface{}](false)

	moveBuilders []MoveBuilder[interface{}]
)

func init() {
	moveBuilders = make([]MoveBuilder[interface{}], 13) // 13 = number of pieces (including BLOCKER)
	moveBuilders[6] = whiteKingMoveBuilder              // WHITE_KING
	moveBuilders[5] = whiteQueenMoveBuilder             // WHITE_QUEEN
	moveBuilders[4] = whiteRookMoveBuilder              // WHITE_ROOK
	moveBuilders[3] = whiteBishopMoveBuilder            // WHITE_BISHOP
	moveBuilders[2] = whiteKnightMoveBuilder            // WHITE_KNIGHT
	moveBuilders[1] = whitePawnMoveBuilder              // WHITE_PAWN
	moveBuilders[12] = blackKingMoveBuilder             // BLACK_KING
	moveBuilders[11] = blackQueenMoveBuilder            // BLACK_QUEEN
	moveBuilders[10] = blackRookMoveBuilder             // BLACK_ROOK
	moveBuilders[9] = blackBishopMoveBuilder            // BLACK_BISHOP
	moveBuilders[8] = blackKnightMoveBuilder            // BLACK_KNIGHT
	moveBuilders[7] = blackPawnMoveBuilder              // BLACK_PAWN
	// BLOCKER (0) left as nil
}

// GetMoveBuilder returns the generic MoveBuilder[interface{}] for a given Piece.
func GetMoveBuilder(piece *Piece) MoveBuilder[interface{}] {
	// The ordinal in Java is the index in the enum, which matches the order in Piece.go
	// We use the same order for moveBuilders.
	switch piece {
	case nil:
		return nil
	}
	// Find index by matching code and color
	switch piece.Code {
	case 'K':
		return moveBuilders[6]
	case 'Q':
		return moveBuilders[5]
	case 'R':
		return moveBuilders[4]
	case 'B':
		return moveBuilders[3]
	case 'N':
		return moveBuilders[2]
	case 'P':
		return moveBuilders[1]
	case 'k':
		return moveBuilders[12]
	case 'q':
		return moveBuilders[11]
	case 'r':
		return moveBuilders[10]
	case 'b':
		return moveBuilders[9]
	case 'n':
		return moveBuilders[8]
	case 'p':
		return moveBuilders[7]
	default:
		return nil
	}
}

// --- MoveBuilder implementations using generics ---

type BasicMoveBuilder[T any] struct {
	deltas  []int
	isWhite bool
}

func NewBasicMoveBuilder[T any](deltas []int, isWhite bool) BasicMoveBuilder[T] {
	return BasicMoveBuilder[T]{deltas: deltas, isWhite: isWhite}
}

func (b BasicMoveBuilder[T]) Build(moves *[]T, board Explorable, from int, moveBuilder MoveConstructor[T]) {
	for _, delta := range b.deltas {
		to := from + delta
		piece := board.GetCapturable(to)
		if piece == nil || piece.CanBeCapturedBy(b.isWhite) {
			*moves = append(*moves, moveBuilder.Create(from, to))
		}
	}
}

type KingMoveBuilder[T any] struct {
	BasicMoveBuilder[T]
}

func NewKingMoveBuilder[T any](isWhite bool) KingMoveBuilder[T] {
	deltas := []int{
		NORTH.Delta(), SOUTH.Delta(), EAST.Delta(), WEST.Delta(),
		NORTH_EAST.Delta(), NORTH_WEST.Delta(), SOUTH_EAST.Delta(), SOUTH_WEST.Delta(),
	}
	return KingMoveBuilder[T]{NewBasicMoveBuilder[T](deltas, isWhite)}
}

type KnightMoveBuilder[T any] struct {
	BasicMoveBuilder[T]
}

func NewKnightMoveBuilder[T any](isWhite bool) KnightMoveBuilder[T] {
	deltas := []int{
		2*NORTH.Delta() + EAST.Delta(),
		2*NORTH.Delta() + WEST.Delta(),
		2*SOUTH.Delta() + EAST.Delta(),
		2*SOUTH.Delta() + WEST.Delta(),
		NORTH.Delta() + 2*EAST.Delta(),
		NORTH.Delta() + 2*WEST.Delta(),
		SOUTH.Delta() + 2*EAST.Delta(),
		SOUTH.Delta() + 2*WEST.Delta(),
	}
	return KnightMoveBuilder[T]{NewBasicMoveBuilder[T](deltas, isWhite)}
}

type PawnMoveBuilder[T any] struct {
	isWhite        bool
	advanceDelta   int
	captureDeltaW  int
	captureDeltaE  int
	twoAdvanceRank int
}

func NewPawnMoveBuilder[T any](isWhite bool) PawnMoveBuilder[T] {
	advanceDelta := NORTH.Delta()
	captureDeltaW := NORTH_WEST.Delta()
	captureDeltaE := NORTH_EAST.Delta()
	twoAdvanceRank := 1
	if !isWhite {
		advanceDelta = SOUTH.Delta()
		captureDeltaW = SOUTH_WEST.Delta()
		captureDeltaE = SOUTH_EAST.Delta()
		twoAdvanceRank = 6
	}
	return PawnMoveBuilder[T]{
		isWhite:        isWhite,
		advanceDelta:   advanceDelta,
		captureDeltaW:  captureDeltaW,
		captureDeltaE:  captureDeltaE,
		twoAdvanceRank: twoAdvanceRank,
	}
}

func (p PawnMoveBuilder[T]) Build(moves *[]T, board Explorable, from int, moveBuilder MoveConstructor[T]) {
	fmt.Printf("TRACE: PawnMoveBuilder.Build called (isWhite=%v, from=%d)\n", p.isWhite, from)
	to := from + p.advanceDelta
	if board.GetCapturable(to) == nil {
		*moves = append(*moves, moveBuilder.Create(from, to))
		to2 := to + p.advanceDelta
		if p.twoAdvanceRank == board.GetRank(from) && board.GetCapturable(to2) == nil {
			*moves = append(*moves, moveBuilder.Create(from, to2))
		}
	}
	toW := from + p.captureDeltaW
	captured := board.GetCapturable(toW)
	if captured != nil && captured.CanBeCapturedBy(p.isWhite) {
		*moves = append(*moves, moveBuilder.Create(from, toW))
	}
	toE := from + p.captureDeltaE
	captured = board.GetCapturable(toE)
	if captured != nil && captured.CanBeCapturedBy(p.isWhite) {
		*moves = append(*moves, moveBuilder.Create(from, toE))
	}
}

type SliderMoveBuilder[T any] struct {
	deltas  []int
	isWhite bool
}

func NewSliderMoveBuilder[T any](directions []Direction, isWhite bool) SliderMoveBuilder[T] {
	deltas := make([]int, len(directions))
	for i, d := range directions {
		deltas[i] = d.Delta()
	}
	return SliderMoveBuilder[T]{deltas: deltas, isWhite: isWhite}
}

func (s SliderMoveBuilder[T]) scanDirection(moves *[]T, board Explorable, from int, delta int, moveBuilder MoveConstructor[T]) {
	to := from + delta
	for {
		piece := board.GetCapturable(to)
		if piece == nil {
			*moves = append(*moves, moveBuilder.Create(from, to))
		} else {
			if piece.CanBeCapturedBy(s.isWhite) {
				*moves = append(*moves, moveBuilder.Create(from, to))
			}
			break
		}
		to += delta
	}
}

func (s SliderMoveBuilder[T]) Build(moves *[]T, board Explorable, from int, moveBuilder MoveConstructor[T]) {
	for _, delta := range s.deltas {
		s.scanDirection(moves, board, from, delta, moveBuilder)
	}
}
