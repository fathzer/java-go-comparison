package chess

import (
	"fmt"
)

// Move represents a chess move from one square to another (by index).
type Move struct {
	from int
	to   int
}

// NewMove creates a Move from two square indices.
func NewMove(from, to int) Move {
	return Move{from: from, to: to}
}

// MoveFromUCI parses a move from UCI notation (e.g. "e2e4").
func MoveFromUCI(uci string) (Move, error) {
	if len(uci) != 4 {
		return Move{}, fmt.Errorf("invalid UCI move: %s", uci)
	}
	from := getSquare(uci[:2])
	to := getSquare(uci[2:])
	return Move{from: from, to: to}, nil
}

// NewMoveFromStrings creates a Move from two square strings (e.g. "e2", "e4").
func NewMoveFromStrings(fromStr, toStr string) (Move, error) {
	from := getSquare(fromStr)
	to := getSquare(toStr)
	return Move{from: from, to: to}, nil
}

// From returns the source square index.
func (m Move) From() int {
	return m.from
}

// To returns the destination square index.
func (m Move) To() int {
	return m.to
}

// Equals checks if two moves are equal.
func (m Move) Equals(other Move) bool {
	return m.from == other.from && m.to == other.to
}

// String returns the move in UCI notation.
func (m Move) String() string {
	// Assuming GetUCI is a function that converts square index to UCI notation.
	return getUCI(m.from) + getUCI(m.to)
}
