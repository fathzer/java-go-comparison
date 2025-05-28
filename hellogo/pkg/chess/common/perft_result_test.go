package chesscommon

import (
	"testing"
)

func TestNewPerftResult(t *testing.T) {
	result := NewPerftResult()
	if result.LeafNodesCount() != 0 {
		t.Errorf("Expected default leafNodesCount to be 0, got %d", result.LeafNodesCount())
	}
	if result.SearchedNodesCount() != 0 {
		t.Errorf("Expected default searchedNodesCount to be 0, got %d", result.SearchedNodesCount())
	}
}

func TestSettersAndGetters(t *testing.T) {
	result := NewPerftResult()

	// Test SetLeafNodesCount and LeafNodesCount
	result.SetLeafNodesCount(42)
	if got := result.LeafNodesCount(); got != 42 {
		t.Errorf("Expected leafNodesCount to be 42, got %d", got)
	}

	// Test IncrementSearchedNodesCount and SearchedNodesCount
	result.IncrementSearchedNodesCount()
	if got := result.SearchedNodesCount(); got != 1 {
		t.Errorf("Expected searchedNodesCount to be 1 after increment, got %d", got)
	}

	result.IncrementSearchedNodesCount()
	if got := result.SearchedNodesCount(); got != 2 {
		t.Errorf("Expected searchedNodesCount to be 2 after second increment, got %d", got)
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name     string
		a        *PerftResult
		b        *PerftResult
		expected bool
	}{
		{"both nil", nil, nil, true},
		{"a nil", nil, NewPerftResult(), false},
		{"b nil", NewPerftResult(), nil, false},
		{"both default", NewPerftResult(), NewPerftResult(), true},
		{"same values", createPerftResult(10, 20), createPerftResult(10, 20), true},
		{"different leaf nodes", createPerftResult(10, 20), createPerftResult(11, 20), false},
		{"different searched nodes", createPerftResult(10, 20), createPerftResult(10, 21), false},
		{"different both", createPerftResult(10, 20), createPerftResult(11, 21), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result bool
			if tt.a == nil {
				result = (tt.b == nil) == tt.expected
			} else {
				result = tt.a.Equals(tt.b) == tt.expected
			}

			if !result {
				t.Errorf("Equals() failed for %s, expected %v", tt.name, tt.expected)
			}

			// Test symmetry
			if tt.a != nil && tt.b != nil {
				if tt.a.Equals(tt.b) != tt.b.Equals(tt.a) {
					t.Errorf("Equals() is not symmetric for %s", tt.name)
				}
			}
		})
	}
}

func TestHashCode(t *testing.T) {
	tests := []struct {
		a        *PerftResult
		b        *PerftResult
		sameHash bool
	}{
		{NewPerftResult(), NewPerftResult(), true},
		{createPerftResult(10, 20), createPerftResult(10, 20), true},
		{createPerftResult(10, 20), createPerftResult(11, 20), false},
		{createPerftResult(10, 20), createPerftResult(10, 21), false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			hashA := tt.a.HashCode()
			hashB := tt.b.HashCode()

			if tt.sameHash && hashA != hashB {
				t.Errorf("Expected equal hash codes, got %d and %d", hashA, hashB)
			} else if !tt.sameHash && hashA == hashB {
				// This is not strictly required (hash collisions can happen),
				// but it's good to test for common cases
				t.Logf("Different objects have same hash code: %d and %d", hashA, hashB)
			}
		})
	}
}

func TestEqualsAndHashCodeConsistency(t *testing.T) {
	r1 := createPerftResult(10, 20)
	r2 := createPerftResult(10, 20)

	if r1.Equals(r2) && r1.HashCode() != r2.HashCode() {
		t.Error("Equal objects must have equal hash codes")
	}
}

// Helper function to create a PerftResult with specific values
func createPerftResult(leafNodes, searchedNodes int64) *PerftResult {
	result := NewPerftResult()
	result.SetLeafNodesCount(leafNodes)
	for i := int64(0); i < searchedNodes; i++ {
		result.IncrementSearchedNodesCount()
	}
	return result
}
