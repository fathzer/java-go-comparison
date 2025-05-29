// Package optimized provides optimized data structures for chess engine implementation.
package optimized

// IntMove provides utility functions for encoding and decoding move information into a single int.
// The encoding uses the following bit layout:
// - from (7 bits): bits 0-6 (0-127)
// - to (7 bits): bits 7-13 (0-127)
// - capture (5 bits): bits 14-18 (0-31)
const (
	fromMask    = 0x7F    // 7 bits for from (0-127)
	toMask      = 0x3F80  // 7 bits for to (shifted left by 7)
	captureMask = 0x7C000 // 5 bits for capture (shifted left by 14)

	toShift      = 7
	captureShift = 14
)

// FromUCI creates a packed move from a UCI string and the current board state.
// The UCI string should be 4 characters long (e.g., "e2e4").
func FromUCI(board *Board, uci string) int {
	if len(uci) != 4 {
		panic("Invalid UCI move: " + uci)
	}
	toSquare := GetSquare(uci[2:])
	capturedPiece := board.getPiece(toSquare)
	return Move(GetSquare(uci[:2]), toSquare, capturedPiece)
}

// Move creates a packed move integer from individual components.
// from: The source square (0-127)
// to: The target square (0-127)
// capture: The capture value (0-31)
// Returns a packed integer containing all three values.
func Move(from, to, capture int) int {
	return (from) |
		(to << toShift) |
		(capture << captureShift)
}

// From extracts the source square from a packed move.
// Returns the source square (0-127).
func From(move int) int {
	return move & fromMask
}

// To extracts the target square from a packed move.
// Returns the target square (0-127).
func To(move int) int {
	return (move & toMask) >> toShift
}

// Capture extracts the capture value from a packed move.
// Returns the capture value (0-31).
func Capture(move int) int {
	return (move & captureMask) >> captureShift
}

// String returns a string representation of the move in UCI format.
// For captures, it also includes the captured piece in parentheses.
func String(move int) string {
	capture := Capture(move)
	if capture == 0 {
		return GetUCI(From(move)) + GetUCI(To(move))
	}
	return GetUCI(From(move)) + "x" + GetUCI(To(move)) + "(" + string(GetCode(capture)) + ")"
}
