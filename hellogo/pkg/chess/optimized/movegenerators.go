// Package optimized provides optimized data structures for chess engine implementation.
package optimized

// MoveGenerators provides move generation for different chess pieces.
type MoveGenerators struct{}

// Direction represents a direction of movement on the chess board.
type direction struct {
	delta int
}

// Constants for directions
var (
	north     = direction{delta: rowWidth}
	south     = direction{delta: -rowWidth}
	east      = direction{delta: +1}
	west      = direction{delta: -1}
	northEast = direction{delta: rowWidth + 1}
	northWest = direction{delta: rowWidth - 1}
	southEast = direction{delta: -rowWidth + 1}
	southWest = direction{delta: -rowWidth - 1}
)

const rowWidth = 10

var (
	whiteKingMoveBuilder   = newKingMoveBuilder(true)
	whiteQueenMoveBuilder  = newSliderMoveBuilder([]direction{north, south, east, west, northEast, northWest, southEast, southWest}, true)
	whiteRookMoveBuilder   = newSliderMoveBuilder([]direction{north, south, east, west}, true)
	whiteBishopMoveBuilder = newSliderMoveBuilder([]direction{northEast, northWest, southEast, southWest}, true)
	whiteKnightMoveBuilder = newKnightMoveBuilder(true)
	whitePawnMoveBuilder   = newPawnMoveBuilder(true)

	blackKingMoveBuilder   = newKingMoveBuilder(false)
	blackQueenMoveBuilder  = newSliderMoveBuilder([]direction{north, south, east, west, northEast, northWest, southEast, southWest}, false)
	blackRookMoveBuilder   = newSliderMoveBuilder([]direction{north, south, east, west}, false)
	blackBishopMoveBuilder = newSliderMoveBuilder([]direction{northEast, northWest, southEast, southWest}, false)
	blackKnightMoveBuilder = newKnightMoveBuilder(false)
	blackPawnMoveBuilder   = newPawnMoveBuilder(false)

	moveBuilders []MoveBuilder
)

func init() {
	moveBuilders = make([]MoveBuilder, BlackKing+1)
	moveBuilders[WhiteKing] = whiteKingMoveBuilder
	moveBuilders[WhiteQueen] = whiteQueenMoveBuilder
	moveBuilders[WhiteRook] = whiteRookMoveBuilder
	moveBuilders[WhiteBishop] = whiteBishopMoveBuilder
	moveBuilders[WhiteKnight] = whiteKnightMoveBuilder
	moveBuilders[WhitePawn] = whitePawnMoveBuilder
	moveBuilders[BlackKing] = blackKingMoveBuilder
	moveBuilders[BlackQueen] = blackQueenMoveBuilder
	moveBuilders[BlackRook] = blackRookMoveBuilder
	moveBuilders[BlackBishop] = blackBishopMoveBuilder
	moveBuilders[BlackKnight] = blackKnightMoveBuilder
	moveBuilders[BlackPawn] = blackPawnMoveBuilder
}

// Get returns the move builder for the given piece.
func (mg *MoveGenerators) Get(piece int) MoveBuilder {
	return moveBuilders[piece]
}

// Get returns the move builder for the given piece.
// This is a package-level function that provides a more convenient way to access move builders.
func Get(piece int) MoveBuilder {
	return moveBuilders[piece]
}

type basicMoveBuilder struct {
	deltas  []int
	isWhite bool
}

func newBasicMoveBuilder(deltas []int, isWhite bool) *basicMoveBuilder {
	return &basicMoveBuilder{
		deltas:  deltas,
		isWhite: isWhite,
	}
}

func (b *basicMoveBuilder) Build(moves *IntList, board *Board, from int) {
	for _, delta := range b.deltas {
		to := from + delta
		piece := board.getPiece(to)
		if piece == None || canBeCapturedBy(piece, b.isWhite) {
			moves.Add(Move(from, to, piece))
		}
	}
}

type kingMoveBuilder struct {
	*basicMoveBuilder
}

func newKingMoveBuilder(isWhite bool) *kingMoveBuilder {
	deltas := []int{
		north.delta, south.delta, east.delta, west.delta,
		northEast.delta, northWest.delta, southEast.delta, southWest.delta,
	}
	return &kingMoveBuilder{
		basicMoveBuilder: newBasicMoveBuilder(deltas, isWhite),
	}
}

type knightMoveBuilder struct {
	*basicMoveBuilder
}

func newKnightMoveBuilder(isWhite bool) *knightMoveBuilder {
	return &knightMoveBuilder{
		basicMoveBuilder: newBasicMoveBuilder(getKnightDeltas(), isWhite),
	}
}

func getKnightDeltas() []int {
	return []int{
		2*north.delta + east.delta,
		2*north.delta + west.delta,
		2*south.delta + east.delta,
		2*south.delta + west.delta,
		north.delta + 2*east.delta,
		north.delta + 2*west.delta,
		south.delta + 2*east.delta,
		south.delta + 2*west.delta,
	}
}

// PawnMoveBuilder generates moves for pawns.
// WARNING: This is a very basic implementation. It does not manage en passant, promotion.
type pawnMoveBuilder struct {
	isWhite        bool
	advanceDelta   int
	captureDeltaWest int
	captureDeltaEast int
	twoAdvanceRank  int
}

func newPawnMoveBuilder(isWhite bool) *pawnMoveBuilder {
	advanceDelta := north.delta
	captureDeltaWest := northWest.delta
	captureDeltaEast := northEast.delta
	twoAdvanceRank := 1
	if !isWhite {
		advanceDelta = south.delta
		captureDeltaWest = southWest.delta
		captureDeltaEast = southEast.delta
		twoAdvanceRank = 6
	}

	return &pawnMoveBuilder{
		isWhite:         isWhite,
		advanceDelta:    advanceDelta,
		captureDeltaWest: captureDeltaWest,
		captureDeltaEast: captureDeltaEast,
		twoAdvanceRank:   twoAdvanceRank,
	}
}

func (p *pawnMoveBuilder) Build(moves *IntList, board *Board, from int) {
	// Single square advance
	to := from + p.advanceDelta
	if board.getPiece(to) == None {
		moves.Add(Move(from, to, None))
		// Two square advance from starting position
		to += p.advanceDelta
		if GetRank(from) == p.twoAdvanceRank && board.getPiece(to) == None {
			moves.Add(Move(from, to, None))
		}
	}

	// Captures
	to = from + p.captureDeltaWest
	captured := board.getPiece(to)
	if captured != None && canBeCapturedBy(captured, p.isWhite) {
		moves.Add(Move(from, to, captured))
	}

	to = from + p.captureDeltaEast
	captured = board.getPiece(to)
	if captured != None && canBeCapturedBy(captured, p.isWhite) {
		moves.Add(Move(from, to, captured))
	}
}

type sliderMoveBuilder struct {
	deltas  []int
	isWhite bool
}

func newSliderMoveBuilder(directions []direction, isWhite bool) *sliderMoveBuilder {
	deltas := make([]int, len(directions))
	for i, d := range directions {
		deltas[i] = d.delta
	}
	return &sliderMoveBuilder{
		deltas:  deltas,
		isWhite: isWhite,
	}
}

func (s *sliderMoveBuilder) scanDirection(moves *IntList, board *Board, from, delta int) {
	to := from + delta
	for {
		piece := board.getPiece(to)
		if piece == None {
			moves.Add(Move(from, to, None))
		} else {
			if canBeCapturedBy(piece, s.isWhite) {
				moves.Add(Move(from, to, piece))
			}
			break
		}
		to += delta
	}
}

func (s *sliderMoveBuilder) Build(moves *IntList, board *Board, from int) {
	for _, delta := range s.deltas {
		s.scanDirection(moves, board, from, delta)
	}
}

// canBeCapturedBy returns true if the piece can be captured by a piece of the given color.
func canBeCapturedBy(piece int, isWhite bool) bool {
	return isWhite != (piece < BlackPawn)
}
