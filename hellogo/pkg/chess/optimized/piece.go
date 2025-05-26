// Package optimized provides optimized data structures for chess engine implementation.
package optimized

// Piece constants
const (
	None = iota
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
	Blocker // Special value for board edges (must be greater than any other piece)
)

// codes maps piece constants to their corresponding FEN characters
var codes = []rune{' ', 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k', 'X'}

// codeToPiece maps FEN characters to piece constants
var codeToPiece = map[rune]int{
	'P': WhitePawn,
	'N': WhiteKnight,
	'B': WhiteBishop,
	'R': WhiteRook,
	'Q': WhiteQueen,
	'K': WhiteKing,
	'p': BlackPawn,
	'n': BlackKnight,
	'b': BlackBishop,
	'r': BlackRook,
	'q': BlackQueen,
	'k': BlackKing,
}

// FromCode converts a FEN character to a piece constant.
// Returns None if the code doesn't correspond to any piece.
func FromCode(code rune) int {
	if piece, ok := codeToPiece[code]; ok {
		return piece
	}
	return None
}

// GetCode returns the FEN character for a piece
func GetCode(piece int) rune {
	return codes[piece]
}

// IsWhite returns true if the piece is white
func IsWhite(piece int) bool {
	return piece <= WhiteKing
}

// CanBeCapturedBy returns true if the piece can be captured by the specified color
func CanBeCapturedBy(piece int, white bool) bool {
	return piece != Blocker && IsWhite(piece) != white
}
