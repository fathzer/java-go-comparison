// Package optimized provides optimized data structures for chess engine implementation.
package optimized

// MoveBuilder is an interface that defines a method for generating legal moves
// from a specific square on a chess board.
//
// Implementations of this interface are responsible for scanning the board
// and adding legal moves to the provided list.
type MoveBuilder interface {
	// Build scans the board for legal moves from a given square and adds them
	// to the provided moves list.
	//
	// moves: The list to which legal moves should be added
	// board: The chess board to explore
	// from: The square (0-127) from which to generate moves
	Build(moves *IntList, board *Board, from int)
}
