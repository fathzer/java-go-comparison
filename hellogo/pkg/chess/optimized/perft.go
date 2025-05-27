package optimized

import (
	chesscommon "hellogo/pkg/chess/common"
)

// Perft performs a Performance Test that walks the move generation tree of strictly legal moves
// to count all the leaf nodes of a certain depth, which can be compared to predetermined values
// and used to isolate bugs.
type Perft struct{}

// NewPerft creates a new Perft instance
func NewPerft() *Perft {
	return &Perft{}
}

// Perft performs a non-bulk Perft (Performance Test) calculation.
// board: The board to run the performance test on.
// depth: The depth to run the performance test to
// whitePlaying: true if white is playing, false otherwise
// Returns a non-null result
func (p *Perft) Perft(board *Board, depth int, whitePlaying bool) *chesscommon.PerftResult {
	return p.perftWithType(board, depth, chesscommon.NonBulk, whitePlaying)
}

// PerftWithType performs a Perft (Performance Test) calculation.
// board: The board to run the performance test on.
// depth: The depth to run the performance test to
// perftType: The type of Perft to run
// whitePlaying: true if white is playing, false otherwise
// Returns a non-null result
func (p *Perft) PerftWithType(board *Board, depth int, perftType chesscommon.PerftType, whitePlaying bool) *chesscommon.PerftResult {
	return p.perftWithType(board, depth, perftType, whitePlaying)
}

func (p *Perft) perftWithType(board *Board, depth int, perftType chesscommon.PerftType, whitePlaying bool) *chesscommon.PerftResult {
	if board == nil {
		panic("Board cannot be nil")
	}
	if depth <= 0 {
		panic("Depth must be greater than 0")
	}

	result := chesscommon.NewPerftResult()
	moveListCache := make([]*IntList, depth+1)
	for i := range moveListCache {
		moveListCache[i] = NewIntList()
	}

	leafNodes := p.perft(moveListCache, board, result, depth, depth, perftType, whitePlaying)
	result.SetLeafNodesCount(leafNodes)
	return result
}

func (p *Perft) perft(moveListCache []*IntList, board *Board, result *chesscommon.PerftResult, depth, originalDepth int, perftType chesscommon.PerftType, whitePlaying bool) int64 {
	result.IncrementSearchedNodesCount()
	moves := moveListCache[depth]
	board.GetMoves(moves, whitePlaying)

	if depth == 1 && perftType == chesscommon.NonBulk {
		return int64(moves.Size())
	} else if depth == 0 {
		return 1
	}

	var leafNodesCount int64
	for i := 0; i < moves.Size(); i++ {
		move := moves.Get(i)
		board.MakeMove(move)
		moveCount := p.perft(moveListCache, board, result, depth-1, originalDepth, perftType, !whitePlaying)
		leafNodesCount += moveCount
		board.UnmakeMove(move)
	}
	return leafNodesCount
}
