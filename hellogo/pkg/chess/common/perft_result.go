// Package chesscommon contains shared types and utilities for the chess engine.
package chesscommon

// PerftResult holds the results of a Perft (Performance Test) calculation.
type PerftResult[T comparable] struct {
	searchedNodesCount int64
	leafNodesCount     int64
	nodesPerMove       map[T]int64
}

// NewPerftResult creates a new PerftResult instance.
func NewPerftResult[T comparable]() *PerftResult[T] {
	return &PerftResult[T]{
		nodesPerMove: make(map[T]int64),
	}
}

// LeafNodesCount returns the number of leaf nodes.
func (p *PerftResult[T]) LeafNodesCount() int64 {
	return p.leafNodesCount
}

// SearchedNodesCount returns the number of nodes for which the move generation has been searched.
func (p *PerftResult[T]) SearchedNodesCount() int64 {
	return p.searchedNodesCount
}

// Divide returns a map of moves to the number of nodes at first depth.
func (p *PerftResult[T]) Divide() map[T]int64 {
	// Return a copy to prevent external modifications
	result := make(map[T]int64, len(p.nodesPerMove))
	for k, v := range p.nodesPerMove {
		result[k] = v
	}
	return result
}

// IncrementSearchedNodesCount increments the searched nodes count by 1.
func (p *PerftResult[T]) IncrementSearchedNodesCount() {
	p.searchedNodesCount++
}

// SetLeafNodesCount sets the leaf nodes count.
func (p *PerftResult[T]) SetLeafNodesCount(count int64) {
	p.leafNodesCount = count
}

// SetNodesPerMove sets the number of nodes for a specific move.
func (p *PerftResult[T]) SetNodesPerMove(move T, nodes int64) {
	p.nodesPerMove[move] = nodes
}
