package optimized

import (
	"testing"
)

func TestPerft(t *testing.T) {
	board := NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	perft := NewPerft()

	// Test invalid inputs
	assertPanics(t, func() { perft.Perft(nil, 1, false) }, "Board cannot be nil")
	assertPanics(t, func() { perft.Perft(board, -1, false) }, "Depth must be greater than 0")
	assertPanics(t, func() { perft.Perft(board, 0, false) }, "Depth must be greater than 0")

	// Test valid positions and depths
	result1 := perft.Perft(board, 1, true)
	if result1.LeafNodesCount() != 20 {
		t.Errorf("Expected 20 leaf nodes, got %d", result1.LeafNodesCount())
	}

	result2 := perft.Perft(board, 2, true)
	if result2.LeafNodesCount() != 400 {
		t.Errorf("Expected 400 leaf nodes, got %d", result2.LeafNodesCount())
	}

	board2 := NewBoard("rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR")
	result3 := perft.Perft(board2, 1, false)
	if result3.LeafNodesCount() != 21 {
		t.Errorf("Expected 21 leaf nodes, got %d", result3.LeafNodesCount())
	}

	result4 := perft.Perft(board2, 2, false)
	if result4.LeafNodesCount() != 463 {
		t.Errorf("Expected 463 leaf nodes, got %d", result4.LeafNodesCount())
	}
}

// Helper function to assert that a function panics with a specific message
func assertPanics(t *testing.T, f func(), expectedMessage string) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("Expected panic, but none occurred")
		} else if r != expectedMessage {
			t.Errorf("Expected panic message '%s', got '%v'", expectedMessage, r)
		}
	}()
	f()
}
