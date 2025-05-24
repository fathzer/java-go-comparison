package chess

// Capturable interface from MoveBuilder.go
// ...existing code...

type Piece struct {
	Code    rune
	IsWhite bool
	ordinal int
}

// Equals checks if two pieces are equal.
func (p *Piece) Equals(other *Piece) bool {
	if p == other {
		return true
	}
	if p == nil || other == nil {
		return false
	}
	return p.Code == other.Code && p.IsWhite == other.IsWhite
}

var (
	BLOCKER      = Piece{'X', true, 0}
	WHITE_PAWN   = Piece{'P', true, 1}
	WHITE_KNIGHT = Piece{'N', true, 2}
	WHITE_BISHOP = Piece{'B', true, 3}
	WHITE_ROOK   = Piece{'R', true, 4}
	WHITE_QUEEN  = Piece{'Q', true, 5}
	WHITE_KING   = Piece{'K', true, 6}
	BLACK_PAWN   = Piece{'p', false, 7}
	BLACK_KNIGHT = Piece{'n', false, 8}
	BLACK_BISHOP = Piece{'b', false, 9}
	BLACK_ROOK   = Piece{'r', false, 10}
	BLACK_QUEEN  = Piece{'q', false, 11}
	BLACK_KING   = Piece{'k', false, 12}
)

var (
	allPieces = []Piece{
		BLOCKER,
		WHITE_PAWN, WHITE_KNIGHT, WHITE_BISHOP, WHITE_ROOK, WHITE_QUEEN, WHITE_KING,
		BLACK_PAWN, BLACK_KNIGHT, BLACK_BISHOP, BLACK_ROOK, BLACK_QUEEN, BLACK_KING,
	}
	codeToPiece map[rune]*Piece
)

func init() {
	codeToPiece = make(map[rune]*Piece, len(allPieces))
	for i := range allPieces {
		p := &allPieces[i]
		if p.Code != BLOCKER.Code {
			codeToPiece[p.Code] = p
		}
	}
}

// FromCode returns a pointer to the Piece for a given character code, or nil if not found.
func FromCode(code rune) *Piece {
	if p, ok := codeToPiece[code]; ok {
		return p
	}
	return nil
}

// GetCode returns the character code for the piece.
func (p *Piece) GetCode() rune {
	return p.Code
}

// IsWhitePiece returns true if the piece is white.
func (p *Piece) IsWhitePiece() bool {
	return p.IsWhite
}

// CanBeCapturedBy implements the Capturable interface.
func (p *Piece) CanBeCapturedBy(white bool) bool {
	return p != nil && p != &BLOCKER && white != p.IsWhite
}

// Ordinal returns the ordinal value of the piece, similar to Java's Enum.ordinal().
func (p *Piece) Ordinal() int {
	return p.ordinal
}
