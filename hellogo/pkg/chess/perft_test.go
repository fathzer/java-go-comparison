package chess

import (
	"testing"
)

func TestPerft(t *testing.T) {
	board, err := NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	if err != nil {
		t.Fatalf("Failed to create board: %v", err)
	}

	perft := NewPerft()

	// Test invalid inputs
	_, err = perft.Perft(nil, 1, false)
	if err == nil {
		t.Error("Expected error for nil board")
	}

	_, err = perft.Perft(board, -1, false)
	if err == nil {
		t.Error("Expected error for negative depth")
	}

	_, err = perft.Perft(board, 0, false)
	if err == nil {
		t.Error("Expected error for zero depth")
	}

	// Test depth 1
	result1, err := perft.Perft(board, 1, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result1.LeafNodes != 20 {
		t.Errorf("Expected 20 leaf nodes, got %d", result1.LeafNodes)
	}
	for _, count := range result1.NodesPerMove {
		if count != 1 {
			t.Errorf("Expected each move to have 1 node, got %d", count)
		}
	}

	// Test depth 2
	result2, err := perft.Perft(board, 2, true)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result2.LeafNodes != 400 {
		t.Errorf("Expected 400 leaf nodes, got %d", result2.LeafNodes)
	}

	// Test with different position
	board2, err := NewBoard("rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR")
	if err != nil {
		t.Fatalf("Failed to create board: %v", err)
	}

	result3, err := perft.Perft(board2, 1, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result3.LeafNodes != 21 {
		t.Errorf("Expected 21 leaf nodes, got %d", result3.LeafNodes)
	}

	result4, err := perft.Perft(board2, 2, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result4.LeafNodes != 463 {
		t.Errorf("Expected 463 leaf nodes, got %d", result4.LeafNodes)
	}
}
