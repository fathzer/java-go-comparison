package chess

import (
	"fmt"

	chesscommon "hellogo/pkg/chess/common"
)

// Perft performs performance testing on move generation.
type Perft struct{}

// NewPerft creates a new Perft instance.
func NewPerft() *Perft {
	return &Perft{}
}

// Perft performs a non-bulk Perft (Performance Test) calculation.
// Perft performs a non-bulk Perft (Performance Test) calculation.
func (p *Perft) Perft(board *Board, depth int, whitePlaying bool) (*chesscommon.PerftResult, error) {
	return p.PerftWithType(board, depth, chesscommon.NonBulk, whitePlaying)
}

// PerftWithType performs a Perft (Performance Test) calculation with the specified type.
func (p *Perft) PerftWithType(board *Board, depth int, ptype chesscommon.PerftType, whitePlaying bool) (*chesscommon.PerftResult, error) {
	if board == nil {
		return nil, fmt.Errorf("board cannot be nil")
	}
	if depth <= 0 {
		return nil, fmt.Errorf("depth must be greater than 0")
	}

	result := chesscommon.NewPerftResult()

	result.SetLeafNodesCount(p.perft(board, result, depth, depth, ptype, whitePlaying))
	return result, nil
}

// perft is the internal recursive function that performs the Perft calculation.
func (p *Perft) perft(board *Board, result *chesscommon.PerftResult, depth, originalDepth int, ptype chesscommon.PerftType, whitePlaying bool) int64 {
	result.IncrementSearchedNodesCount()

	moves := board.GetMoves(whitePlaying)

	if depth == 1 && ptype == chesscommon.NonBulk {
		return int64(len(moves))
	} else if depth == 0 {
		return 1
	}

	var leafNodes int64
	for _, move := range moves {
		if err := board.MakeMove(&move); err != nil {
			// Skip invalid moves
			continue
		}

		moveCount := p.perft(board, result, depth-1, originalDepth, ptype, !whitePlaying)
		leafNodes += moveCount

		// Unmake the move to restore the board state
		if err := board.UnmakeMove(); err != nil {
			// This should not happen with legal moves
			panic(fmt.Sprintf("failed to unmake move: %v", err))
		}
	}

	return leafNodes
}
