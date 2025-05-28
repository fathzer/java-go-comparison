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

// Equals checks if two PerftResult instances are equal.
func (p *PerftResult) Equals(other *PerftResult) bool {
	if other == nil {
		return false
	}
	return p.searchedNodesCount == other.searchedNodesCount && p.leafNodesCount == other.leafNodesCount
}

// HashCode returns a hash code for the PerftResult.
// This implementation follows the same logic as Java's Objects.hash()
func (p *PerftResult) HashCode() int {
	const prime = 31
	result := int(1)
	result = prime*result + int(p.searchedNodesCount^(p.searchedNodesCount>>32))
	result = prime*result + int(p.leafNodesCount^(p.leafNodesCount>>32))
	return result
}
