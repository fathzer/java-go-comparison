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
	from, err := GetSquare(uci[:2])
	if err != nil {
		return Move{}, err
	}
	to, err := GetSquare(uci[2:])
	if err != nil {
		return Move{}, err
	}
	return Move{from: from, to: to}, nil
}

// NewMoveFromStrings creates a Move from two square strings (e.g. "e2", "e4").
func NewMoveFromStrings(fromStr, toStr string) (Move, error) {
	from, err := GetSquare(fromStr)
	if err != nil {
		return Move{}, err
	}
	to, err := GetSquare(toStr)
	if err != nil {
		return Move{}, err
	}
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
	return GetUCI(m.from) + GetUCI(m.to)
}

// --- Helper functions for square conversion ---
// These are placeholders; implement them to match your Board logic.

func GetSquare(uciSquare string) (int, error) {
	if len(uciSquare) != 2 {
		return 0, fmt.Errorf("invalid UCI square: %s", uciSquare)
	}
	file := uciSquare[0]
	rank := uciSquare[1]
	if rank < '1' || rank > '8' || file < 'a' || file > 'h' {
		return 0, fmt.Errorf("invalid UCI square: %s", uciSquare)
	}
	// 21+ 10*(rank-1) + (file-'a')
	return 21 + 10*int(rank-'1') + int(file-'a'), nil
}

func GetUCI(square int) string {
	sq := square - 21
	file := sq % 10
	rank := sq / 10
	return fmt.Sprintf("%c%d", 'a'+file, rank+1)
}
