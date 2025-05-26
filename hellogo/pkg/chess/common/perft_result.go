// Package chesscommon contains shared types and utilities for the chess engine.
package chesscommon

// PerftResult holds the results of a Perft (Performance Test) calculation.
type PerftResult struct {
	searchedNodesCount int64
	leafNodesCount     int64
}

// NewPerftResult creates a new PerftResult instance.
func NewPerftResult() *PerftResult {
	return &PerftResult{}
}

// LeafNodesCount returns the number of leaf nodes.
func (p *PerftResult) LeafNodesCount() int64 {
	return p.leafNodesCount
}

// SearchedNodesCount returns the number of nodes for which the move generation has been searched.
func (p *PerftResult) SearchedNodesCount() int64 {
	return p.searchedNodesCount
}

// IncrementSearchedNodesCount increments the searched nodes count by 1.
func (p *PerftResult) IncrementSearchedNodesCount() {
	p.searchedNodesCount++
}

// SetLeafNodesCount sets the leaf nodes count.
func (p *PerftResult) SetLeafNodesCount(count int64) {
	p.leafNodesCount = count
}
