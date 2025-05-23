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
	whiteKingMoveBuilder   = &kingMoveBuilder{isWhite: true}
	whiteQueenMoveBuilder  = newSliderMoveBuilder([]Direction{north, south, east, west, northEast, northWest, southEast, southWest}, true)
	whiteRookMoveBuilder   = newSliderMoveBuilder([]Direction{north, south, east, west}, true)
	whiteBishopMoveBuilder = newSliderMoveBuilder([]Direction{northEast, northWest, southEast, southWest}, true)
	whiteKnightMoveBuilder = &knightMoveBuilder{isWhite: true}
	whitePawnMoveBuilder   = &pawnMoveBuilder{isWhite: true}

	blackKingMoveBuilder   = &kingMoveBuilder{isWhite: false}
	blackQueenMoveBuilder  = newSliderMoveBuilder([]Direction{north, south, east, west, northEast, northWest, southEast, southWest}, false)
	blackRookMoveBuilder   = newSliderMoveBuilder([]Direction{north, south, east, west}, false)
	blackBishopMoveBuilder = newSliderMoveBuilder([]Direction{northEast, northWest, southEast, southWest}, false)
	blackKnightMoveBuilder = &knightMoveBuilder{isWhite: false}
	blackPawnMoveBuilder   = &pawnMoveBuilder{isWhite: false}
)

// GetMoveGenerator returns the appropriate move generator for a given piece.
func GetMoveGenerator(piece *Piece) MoveBuilder {
	if piece == nil {
		return nil
	}

	switch piece.Code {
	case WHITE_KING.Code:
		return whiteKingMoveBuilder
	case BLACK_KING.Code:
		return blackKingMoveBuilder
	case WHITE_QUEEN.Code:
		return whiteQueenMoveBuilder
	case BLACK_QUEEN.Code:
		return blackQueenMoveBuilder
	case WHITE_ROOK.Code:
		return whiteRookMoveBuilder
	case BLACK_ROOK.Code:
		return blackRookMoveBuilder
	case WHITE_BISHOP.Code:
		return whiteBishopMoveBuilder
	case BLACK_BISHOP.Code:
		return blackBishopMoveBuilder
	case WHITE_KNIGHT.Code:
		return whiteKnightMoveBuilder
	case BLACK_KNIGHT.Code:
		return blackKnightMoveBuilder
	case WHITE_PAWN.Code:
		return whitePawnMoveBuilder
	case BLACK_PAWN.Code:
		return blackPawnMoveBuilder
	default:
		return nil
	}
}

type kingMoveBuilder struct {
	isWhite bool
}

func (b *kingMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	directions := []Direction{north, south, east, west, northEast, northWest, southEast, southWest}
	for _, dir := range directions {
		to := from + int(dir)
		piece := board.GetPieceByIndex(to)
		if piece == nil || piece.CanBeCapturedBy(b.isWhite) {
			*moves = append(*moves, Move{from: from, to: to})
		}
	}
}

type knightMoveBuilder struct {
	isWhite bool
}

func (b *knightMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	deltas := []int{
		int(2*north + east),
		int(2*north + west),
		int(2*south + east),
		int(2*south + west),
		int(north + 2*east),
		int(north + 2*west),
		int(south + 2*east),
		int(south + 2*west),
	}

	for _, delta := range deltas {
		to := from + delta
		piece := board.GetPieceByIndex(to)
		if piece == nil || piece.CanBeCapturedBy(b.isWhite) {
			*moves = append(*moves, Move{from: from, to: to})
		}
	}
}

type pawnMoveBuilder struct {
	isWhite        bool
	advanceDelta   int
	twoAdvanceRank int
}

func (b *pawnMoveBuilder) Build(moves *[]Move, board *Board, from int) {
	advanceDelta := int(north)
	if !b.isWhite {
		advanceDelta = int(south)
	}

	// Single advance
	to := from + advanceDelta
	if board.GetPieceByIndex(to) == nil {
		*moves = append(*moves, Move{from: from, to: to})

		// Double advance from starting position
		twoAdvanceRank := 1
		if !b.isWhite {
			twoAdvanceRank = 6
		}
		if GetRank(from) == twoAdvanceRank {
			to += advanceDelta
			if board.GetPieceByIndex(to) == nil {
				*moves = append(*moves, Move{from: from, to: to})
			}
		}
	}

	// Captures
	for _, captureDelta := range []int{int(northWest), int(northEast)} {
		if !b.isWhite {
			captureDelta = -captureDelta
		}
		to = from + captureDelta
		captured := board.GetPieceByIndex(to)
		if captured != nil && captured.CanBeCapturedBy(b.isWhite) {
			*moves = append(*moves, Move{from: from, to: to})
		}
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
		piece := board.GetPieceByIndex(to)
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
