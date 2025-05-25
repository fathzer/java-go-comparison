package chess

import (
	"testing"
)

const (
	errUnexpected = "Unexpected error: %v"
	errCreateBoard = "Failed to create board: %v"
)

func TestPerft(t *testing.T) {
	board, err := NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	if err != nil {
		t.Fatalf(errCreateBoard, err)
	}

	t.Run("Invalid inputs", func(t *testing.T) {
		perft := NewPerft()
		
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
	})


	t.Run("Starting position", func(t *testing.T) {
		perft := NewPerft()

		t.Run("Depth 1", func(t *testing.T) {
			result, err := perft.Perft(board, 1, true)
			if err != nil {
				t.Fatalf(errUnexpected, err)
			}
			if result.LeafNodesCount() != 20 {
				t.Errorf("Expected 20 leaf nodes, got %d", result.LeafNodesCount())
			}
			for _, count := range result.Divide() {
				if count != 1 {
					t.Errorf("Expected each move to have 1 node, got %d", count)
				}
			}
		})

		t.Run("Depth 2", func(t *testing.T) {
			result, err := perft.Perft(board, 2, true)
			if err != nil {
				t.Fatalf(errUnexpected, err)
			}
			if result.LeafNodesCount() != 400 {
				t.Errorf("Expected 400 leaf nodes, got %d", result.LeafNodesCount())
			}
		})
	})

	t.Run("Custom position", func(t *testing.T) {
		board2, err := NewBoard("rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR")
		if err != nil {
			t.Fatalf(errCreateBoard, err)
		}

		perft := NewPerft()

		t.Run("Depth 1", func(t *testing.T) {
			result, err := perft.Perft(board2, 1, false)
			if err != nil {
				t.Fatalf(errUnexpected, err)
			}
			if result.LeafNodesCount() != 21 {
				t.Errorf("Expected 21 leaf nodes, got %d", result.LeafNodesCount())
			}
		})

		t.Run("Depth 2", func(t *testing.T) {
			result, err := perft.Perft(board2, 2, false)
			if err != nil {
				t.Fatalf(errUnexpected, err)
			}
			if result.LeafNodesCount() != 463 {
				t.Errorf("Expected 463 leaf nodes, got %d", result.LeafNodesCount())
			}
		})
	})
}
