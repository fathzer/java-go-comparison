package chess

import (
	"testing"
)

const (
	startingFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
)

func TestFenParsing(t *testing.T) {
	// Standard chess starting position FEN (only piece placement part)
	board, err := NewBoard(startingFEN)
	if err != nil {
		t.Fatalf("Failed to parse FEN: %v", err)
	}

	// Test specific piece positions
	piece := board.GetPiece("d1")
	if piece == nil || !piece.Equals(&WHITE_QUEEN) {
		t.Errorf("Expected white queen at d1, got %v", piece)
	}

	// Test that blockers are set correctly
	testBlocker := func(i int) {
		if board.pieces[i] != &BLOCKER {
			t.Errorf("Expected blocker at index %d, got %v", i, board.pieces[i])
		}
	}

	// Test border blockers
	for i := 0; i < 20; i++ {
		testBlocker(i)
	}
	for i := 100; i < 120; i++ {
		testBlocker(i)
	}
}

func TestCopyConstructor(t *testing.T) {
	original, err := NewBoard(startingFEN)
	if err != nil {
		t.Fatalf("Failed to create board: %v", err)
	}

	copy := NewBoardCopy(original)

	// Verify it's a deep copy
	if copy == original {
		t.Error("Copy should be a different instance")
	}

	// Verify the pieces slice is a different instance
	if &copy.pieces[0] == &original.pieces[0] {
		t.Error("Pieces slice should be a deep copy")
	}

	// Make a move on the copy and verify original is unchanged
	move, _ := NewMoveFromStrings("e2", "e4")
	copy.MakeMove(&move)

	if copy.GetPiece("e4") == nil || copy.GetPiece("e2") != nil {
		t.Error("Move should be applied to copy")
	}

	if original.GetPiece("e4") != nil || !original.GetPiece("e2").Equals(&WHITE_PAWN) {
		t.Error("Original board should be unchanged")
	}
}

func TestInvalidFenThrows(t *testing.T) {
	tests := []struct {
		name string
		fen  string
	}{
		{"badFileCount1", "rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR"},  // 9 is invalid
		{"badFileCount2", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBN"},   // Last row missed one file
		{"badFileCount3", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP2/RNBQKBNR"},  // 7th row has an extra file
		{"badRankCount1", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/"}, // Extra rank
		{"badRankCount2", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP"},           // Missing rank
		{"badPiece", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQXBNR"},       // X is not a valid piece
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBoard(tt.fen)
			if err == nil {
				t.Errorf("Expected error for FEN: %s", tt.fen)
			}
		})
	}
}

func TestMakeMove(t *testing.T) {
	// Start with initial position
	board, _ := NewBoard(startingFEN)

	// Make a pawn move e2-e4
	move, _ := NewMoveFromStrings("e2", "e4")
	err := board.MakeMove(&move)
	if err != nil {
		t.Fatalf("Failed to make move: %v", err)
	}

	// Verify the move was made
	if board.GetPiece("e2") != nil {
		t.Error("Source square should be empty after move")
	}
	if board.GetPiece("e4") == nil || !board.GetPiece("e4").Equals(&WHITE_PAWN) {
		t.Error("Piece should be at destination")
	}
}

func TestMakeCapture(t *testing.T) {
	// Set up a position where white can capture a pawn
	board, _ := NewBoard("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR")

	// Make a capture exd5
	capture, _ := NewMoveFromStrings("e4", "d5")
	err := board.MakeMove(&capture)
	if err != nil {
		t.Fatalf("Failed to make capture: %v", err)
	}

	// Verify the capture
	if board.GetPiece("e4") != nil {
		t.Error("Source square should be empty after capture")
	}
	if board.GetPiece("d5") == nil || !board.GetPiece("d5").Equals(&WHITE_PAWN) {
		t.Error("Capturing piece should be at destination")
	}
}

func TestUnmakeMove(t *testing.T) {
	// Start with initial position
	original, _ := NewBoard(startingFEN)
	board := NewBoardCopy(original)

	// Make a move e2-e4
	move, _ := NewMoveFromStrings("e2", "e4")
	board.MakeMove(&move)

	// Unmake the move
	err := board.UnmakeMove()
	if err != nil {
		t.Fatalf("Failed to unmake move: %v", err)
	}

	// Board should be back to original state
	for rank := '1'; rank <= '8'; rank++ {
		for file := 'a'; file <= 'h'; file++ {
			square := string([]rune{file, rank})
			expected := original.GetPiece(square)
			actual := board.GetPiece(square)
			if expected != actual {
				t.Errorf("Pieces don't match at %s: expected %v, got %v", square, expected, actual)
			}
		}
	}
}

func TestUnmakeCapture(t *testing.T) {
	// Set up a position where white can capture a pawn
	original, _ := NewBoard("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR")
	board := NewBoardCopy(original)

	// Make a capture exd5
	capture, _ := NewMoveFromStrings("e4", "d5")
	board.MakeMove(&capture)

	// Unmake the capture
	err := board.UnmakeMove()
	if err != nil {
		t.Fatalf("Failed to unmake capture: %v", err)
	}

	// Board should be back to original state
	for rank := '1'; rank <= '8'; rank++ {
		for file := 'a'; file <= 'h'; file++ {
			square := string([]rune{file, rank})
			expected := original.GetPiece(square)
			actual := board.GetPiece(square)
			if expected != actual {
				t.Errorf("Pieces don't match at %s: expected %v, got %v", square, expected, actual)
			}
		}
	}
}

func TestUnmakeWithoutMoves(t *testing.T) {
	board, _ := NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
	err := board.UnmakeMove()
	if err == nil {
		t.Error("Expected error when unmaking move with no moves to unmake")
	}
}

func TestGetMoves(t *testing.T) {
	board, _ := NewBoard("8/8/8/8/1k6/8/pK6/Q7")

	// Test white moves
	whiteMoves := board.GetMoves(true)
	expectedWhiteMoves := []Move{}

	// Add queen moves
	for _, to := range []string{"a2", "b1", "c1", "d1", "e1", "f1", "g1", "h1"} {
		m, _ := NewMoveFromStrings("a1", to)
		expectedWhiteMoves = append(expectedWhiteMoves, m)
	}

	// Add king moves
	for _, to := range []string{"a3", "b3", "c3", "a2", "c2", "b1", "c1"} {
		m, _ := NewMoveFromStrings("b2", to)
		expectedWhiteMoves = append(expectedWhiteMoves, m)
	}

	// Check white moves
	if len(whiteMoves) != len(expectedWhiteMoves) {
		t.Errorf("Expected %d white moves, got %d", len(expectedWhiteMoves), len(whiteMoves))
		// Print actual moves for debugging
		for i, m := range whiteMoves {
			t.Logf("Move %d: %s -> %s", i, getUCI(m.From()), getUCI(m.To()))
		}
	}

	// Verify each expected move exists in actual moves
	for _, expected := range expectedWhiteMoves {
		found := false
		for _, actual := range whiteMoves {
			if actual.From() == expected.From() && actual.To() == expected.To() {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected move %s -> %s not found in actual moves",
				getUCI(expected.From()), getUCI(expected.To()))
		}
	}

	// Test black moves
	blackMoves := board.GetMoves(false)
	expectedBlackSquares := map[string]bool{
		"a5": true, "b5": true, "c5": true,
		"a4": true, "c4": true,
		"a3": true, "b3": true, "c3": true,
	}

	if len(blackMoves) != len(expectedBlackSquares) {
		t.Errorf("Expected %d black moves, got %d", len(expectedBlackSquares), len(blackMoves))
		// Print actual moves for debugging
		for i, m := range blackMoves {
			t.Logf("Black move %d: %s -> %s", i, getUCI(m.From()), getUCI(m.To()))
		}
	}

	// Verify each actual move is in expected squares
	for _, move := range blackMoves {
		toSquare := getUCI(move.To())
		if !expectedBlackSquares[toSquare] {
			t.Errorf("Unexpected black move to %s", toSquare)
		}
	}
}
