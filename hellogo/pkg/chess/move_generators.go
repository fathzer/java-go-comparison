package chess

// Direction represents a direction of movement on the board.
type Direction int

const (
	// ROW_WIDTH is the width of a row in the board representation.
	ROW_WIDTH = 10

	north     Direction = ROW_WIDTH
	south     Direction = -ROW_WIDTH
	east      Direction = 1
	west      Direction = -1
	northEast Direction = ROW_WIDTH + 1
	northWest Direction = ROW_WIDTH - 1
	southEast Direction = -ROW_WIDTH + 1
	southWest Direction = -ROW_WIDTH - 1
)

var (
	whiteKingMoveBuilder   = newKingMoveBuilder(true)
	whiteQueenMoveBuilder  = newSliderMoveBuilder([]Direction{north, south, east, west, northEast, northWest, southEast, southWest}, true)
	whiteRookMoveBuilder   = newSliderMoveBuilder([]Direction{north, south, east, west}, true)
	whiteBishopMoveBuilder = newSliderMoveBuilder([]Direction{northEast, northWest, southEast, southWest}, true)
	whiteKnightMoveBuilder = newKnightMoveBuilder(true)
	whitePawnMoveBuilder   = newPawnMoveBuilder(true)

	blackKingMoveBuilder   = newKingMoveBuilder(false)
	blackQueenMoveBuilder  = newSliderMoveBuilder([]Direction{north, south, east, west, northEast, northWest, southEast, southWest}, false)
	blackRookMoveBuilder   = newSliderMoveBuilder([]Direction{north, south, east, west}, false)
	blackBishopMoveBuilder = newSliderMoveBuilder([]Direction{northEast, northWest, southEast, southWest}, false)
	blackKnightMoveBuilder = newKnightMoveBuilder(false)
	blackPawnMoveBuilder   = newPawnMoveBuilder(false)
)

// moveGeneratorMap maps piece codes to their corresponding move builders.
var moveGeneratorMap []MoveBuilder

func init() {
	moveGeneratorMap = make([]MoveBuilder, len(allPieces))

	// Map each piece to its move builder using its ordinal
	moveGeneratorMap[WHITE_KING.Ordinal()] = whiteKingMoveBuilder
	moveGeneratorMap[BLACK_KING.Ordinal()] = blackKingMoveBuilder
	moveGeneratorMap[WHITE_QUEEN.Ordinal()] = whiteQueenMoveBuilder
	moveGeneratorMap[BLACK_QUEEN.Ordinal()] = blackQueenMoveBuilder
	moveGeneratorMap[WHITE_ROOK.Ordinal()] = whiteRookMoveBuilder
	moveGeneratorMap[BLACK_ROOK.Ordinal()] = blackRookMoveBuilder
	moveGeneratorMap[WHITE_BISHOP.Ordinal()] = whiteBishopMoveBuilder
	moveGeneratorMap[BLACK_BISHOP.Ordinal()] = blackBishopMoveBuilder
	moveGeneratorMap[WHITE_KNIGHT.Ordinal()] = whiteKnightMoveBuilder
	moveGeneratorMap[BLACK_KNIGHT.Ordinal()] = blackKnightMoveBuilder
	moveGeneratorMap[WHITE_PAWN.Ordinal()] = whitePawnMoveBuilder
	moveGeneratorMap[BLACK_PAWN.Ordinal()] = blackPawnMoveBuilder
}

// GetMoveGenerator returns the appropriate move generator for a given piece.
// It uses the piece's ordinal for direct array lookup for better performance.
func GetMoveGenerator(piece *Piece) MoveBuilder {
	if piece == nil {
		return nil
	}
	return moveGeneratorMap[piece.Ordinal()]
}

type kingMoveBuilder struct {
	isWhite    bool
	directions []Direction
}

func newKingMoveBuilder(isWhite bool) *kingMoveBuilder {
	return &kingMoveBuilder{
		isWhite:    isWhite,
		directions: []Direction{north, south, east, west, northEast, northWest, southEast, southWest},
	}
}

func (b *kingMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	for _, dir := range b.directions {
		to := from + int(dir)
		piece := board.getPieceByIndex(to)
		if piece == nil || piece.CanBeCapturedBy(b.isWhite) {
			*moves = append(*moves, Move{from: from, to: to})
		}
	}
}

type knightMoveBuilder struct {
	isWhite bool
	deltas  []int
}

func newKnightMoveBuilder(isWhite bool) *knightMoveBuilder {
	return &knightMoveBuilder{
		isWhite: isWhite,
		deltas: []int{
			int(2*north + east),
			int(2*north + west),
			int(2*south + east),
			int(2*south + west),
			int(north + 2*east),
			int(north + 2*west),
			int(south + 2*east),
			int(south + 2*west),
		},
	}
}

func (b *knightMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	for _, delta := range b.deltas {
		to := from + delta
		piece := board.getPieceByIndex(to)
		if piece == nil || piece.CanBeCapturedBy(b.isWhite) {
			*moves = append(*moves, Move{from: from, to: to})
		}
	}
}

type pawnMoveBuilder struct {
	isWhite          bool
	advanceDelta     int
	twoAdvanceRank   int
	captureDeltaEast int
	captureDeltaWest int
}

func newPawnMoveBuilder(isWhite bool) *pawnMoveBuilder {
	advanceDelta := int(north)
	twoAdvanceRank := 1
	captureDeltaEast := int(northEast)
	captureDeltaWest := int(northWest)
	if !isWhite {
		advanceDelta = int(south)
		twoAdvanceRank = 6
		captureDeltaEast = int(southEast)
		captureDeltaWest = int(southWest)
	}
	return &pawnMoveBuilder{
		isWhite:          isWhite,
		advanceDelta:     advanceDelta,
		twoAdvanceRank:   twoAdvanceRank,
		captureDeltaEast: captureDeltaEast,
		captureDeltaWest: captureDeltaWest,
	}
}

func (b *pawnMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	// Single advance
	to := from + b.advanceDelta
	if board.getPieceByIndex(to) == nil {
		*moves = append(*moves, Move{from: from, to: to})

		// Double advance from starting position
		if getRank(from) == b.twoAdvanceRank {
			to += b.advanceDelta
			if board.getPieceByIndex(to) == nil {
				*moves = append(*moves, Move{from: from, to: to})
			}
		}
	}

	// Captures
	to = from + b.captureDeltaWest
	captured := board.getPieceByIndex(to)
	if captured != nil && captured.CanBeCapturedBy(b.isWhite) {
		*moves = append(*moves, Move{from: from, to: to})
	}
	to = from + b.captureDeltaEast
	captured = board.getPieceByIndex(to)
	if captured != nil && captured.CanBeCapturedBy(b.isWhite) {
		*moves = append(*moves, Move{from: from, to: to})
	}
}

type sliderMoveBuilder struct {
	directions []Direction
	isWhite    bool
}

func newSliderMoveBuilder(directions []Direction, isWhite bool) *sliderMoveBuilder {
	return &sliderMoveBuilder{
		directions: directions,
		isWhite:    isWhite,
	}
}

func (b *sliderMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	for _, dir := range b.directions {
		b.scanDirection(moves, board, from, int(dir))
	}
}

func (b *sliderMoveBuilder) scanDirection(moves *[]Move, board *Board, from, delta int) {
	to := from + delta
	for {
		piece := board.getPieceByIndex(to)
		if piece == nil {
			*moves = append(*moves, Move{from: from, to: to})
		} else {
			if piece.CanBeCapturedBy(b.isWhite) {
				*moves = append(*moves, Move{from: from, to: to})
			}
			break
		}
		to += delta
	}
}
