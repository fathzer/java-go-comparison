package chess

import (
	"fmt"
)

// Type represents the type of Perft calculation.
// Please note that since the move generator generates only legal moves, both types should yield the same result.
type PerftType int

const (
	// NonBulk is a non-bulk Perft (Performance Test) calculation; moves at last depth are not played.
	NonBulk PerftType = iota
	// Bulk is a bulk Perft (Performance Test) calculation.
	Bulk
)

// PerftResult holds the results of a Perft (Performance Test) calculation.
type PerftResult struct {
	SearchedNodes int64
	LeafNodes     int64
	NodesPerMove  map[Move]int64
}

// Perft performs performance testing on move generation.
type Perft struct{}

// NewPerft creates a new Perft instance.
func NewPerft() *Perft {
	return &Perft{}
}

// Perft performs a non-bulk Perft (Performance Test) calculation.
func (p *Perft) Perft(board *Board, depth int, whitePlaying bool) (*PerftResult, error) {
	return p.PerftWithType(board, depth, NonBulk, whitePlaying)
}

// PerftWithType performs a Perft (Performance Test) calculation with the specified type.
func (p *Perft) PerftWithType(board *Board, depth int, ptype PerftType, whitePlaying bool) (*PerftResult, error) {
	if board == nil {
		return nil, fmt.Errorf("board cannot be nil")
	}
	if depth <= 0 {
		return nil, fmt.Errorf("depth must be greater than 0")
	}

	result := &PerftResult{
		NodesPerMove: make(map[Move]int64),
	}

	result.LeafNodes = p.perft(board, result, depth, depth, ptype, whitePlaying)
	return result, nil
}

// perft is the internal recursive function that performs the Perft calculation.
func (p *Perft) perft(board *Board, result *PerftResult, depth, originalDepth int, ptype PerftType, whitePlaying bool) int64 {
	result.SearchedNodes++

	moves := board.GetMoves(whitePlaying)

	if depth == 1 && ptype == NonBulk {
		return int64(len(moves))
	} else if depth == 0 {
		return 1
	}

	var leafNodes int64
	for _, move := range moves {
		if err := board.MakeMove(move); err != nil {
			// Skip invalid moves
			continue
		}

		moveCount := p.perft(board, result, depth-1, originalDepth, ptype, !whitePlaying)

		if depth == originalDepth {
			result.NodesPerMove[move] = moveCount
		}

		leafNodes += moveCount

		// Unmake the move to restore the board state
		if err := board.UnmakeMove(); err != nil {
			// This should not happen with legal moves
			panic(fmt.Sprintf("failed to unmake move: %v", err))
		}
	}

	return leafNodes
}

// Divide returns the number of nodes per move at the first depth.
func (r *PerftResult) Divide() map[Move]int64 {
	return r.NodesPerMove
}
