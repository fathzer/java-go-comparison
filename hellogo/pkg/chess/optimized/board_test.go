package optimized

import (
	"testing"
)

const (
	initialPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
)

func TestFenParsing(t *testing.T) {
	// Standard chess starting position FEN (only piece placement part)
	board := NewBoard(initialPosition)
	if board == nil {
		t.Error("Expected board to be created, got nil")
	}
	if piece := board.GetPiece("d1"); piece != WhiteQueen {
		t.Errorf("Expected white queen at d1, got %d", piece)
	}

	// Test that blockers are set correctly
	expectBlocker := func(i int) {
		if piece := board.getPiece(i); piece != Blocker {
			t.Errorf("Expected blocker at %d, got %d", i, piece)
		}
	}
	for i := 0; i < 20; i++ {
		expectBlocker(i)
	}
	for i := 100; i < 120; i++ {
		expectBlocker(i)
	}
}

func TestBoardCopy(t *testing.T) {
	original := NewBoard(initialPosition)
	// Create a copy using the copy constructor
	copy := NewBoardCopy(original)
	if copy == original {
		t.Error("Expected copy to be a different instance than original")
	}
}

func TestInvalidFenThrows(t *testing.T) {
	tests := []struct {
		name string
		fen  string
	}{
		{"badFileCount", "rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR"},  // 9 is invalid
		{"badFileCount2", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBN"},  // Last row missed one file
		{"badFileCount3", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP2/RNBQKBNR"}, // 7th row has an extra file
		{"badRankCount", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/"}, // Last rank is invalid
		{"badRankCount2", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP"},          // Last rank is missing
		{"badPiece", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQXBNR"},      // X is not a valid piece
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for FEN: %s", tt.fen)
				}
			}()
			NewBoard(tt.fen)
		})
	}
}

func TestMakeMove(t *testing.T) {
	// Start with initial position
	board := NewBoard(initialPosition)

	// Make a pawn move e2-e4
	move := FromUCI(board, "e2e4")
	board.MakeMove(move)

	// Verify the move was made
	if piece := board.GetPiece("e2"); piece != None {
		t.Error("Source square should be empty after move")
	}
	if piece := board.GetPiece("e4"); piece != WhitePawn {
		t.Error("Piece should be at destination")
	}
}

func TestMakeCapture(t *testing.T) {
	// Set up a position where white can capture a pawn
	board := NewBoard("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR")

	// Make a capture exd5
	capture := FromUCI(board, "e4d5")
	board.MakeMove(capture)

	// Verify the capture
	if piece := board.GetPiece("e4"); piece != None {
		t.Error("Source square should be empty after capture")
	}
	if piece := board.GetPiece("d5"); piece != WhitePawn {
		t.Error("Capturing piece should be at destination")
	}
}

func TestUnmakeMove(t *testing.T) {
	// Start with initial position
	original := NewBoard(initialPosition)
	// Create a copy by creating a new board from the original's FEN string
	board := NewBoardCopy(original)

	// Make a move
	move := FromUCI(board, "e2e4") // e2-e4
	board.MakeMove(move)

	// Unmake the move
	board.UnmakeMove(move)

	// Board should be back to original state
	for rank := '1'; rank <= '8'; rank++ {
		for file := 'a'; file <= 'h'; file++ {
			square := string(file) + string(rank)
			expected := original.GetPiece(square)
			actual := board.GetPiece(square)
			if expected != actual {
				t.Errorf("Pieces should match at %s: expected %d, got %d", square, expected, actual)
			}
		}
	}
}

func TestUnmakeCapture(t *testing.T) {
	// Set up a position where white can capture a pawn
	fen := "rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR"
	original := NewBoard(fen)
	// Create a copy by creating a new board from the original's FEN string
	board := NewBoardCopy(original)

	// Make a capture
	capture := FromUCI(board, "e4d5") // e4xd5
	board.MakeMove(capture)

	// Unmake the capture
	board.UnmakeMove(capture)

	// Board should be back to original state
	for rank := '1'; rank <= '8'; rank++ {
		for file := 'a'; file <= 'h'; file++ {
			square := string(file) + string(rank)
			expected := original.GetPiece(square)
			actual := board.GetPiece(square)
			if expected != actual {
				t.Errorf("Pieces should match at %s: expected %d, got %d", square, expected, actual)
			}
		}
	}
}

func TestGetMoves(t *testing.T) {
	board := NewBoard("8/8/8/8/1k6/8/pK6/Q7")
	moves := NewIntList()

	// Test white moves
	board.GetMoves(moves, true)
	expected := NewIntList()
	expected.AddAll(parseMoveList(board, "a1", "a2 b1 c1 d1 e1 f1 g1 h1"))
	expected.AddAll(parseMoveList(board, "b2", "a3 b3 c3 a2 c2 b1 c1"))
	testMoves(t, expected, moves)

	// Test black moves
	moves.Clear()
	board.GetMoves(moves, false)
	expectedMoves := parseMoveList(board, "b4", "a5 b5 c5 a4 c4 a3 b3 c3")
	testMoves(t, expectedMoves, moves)
	if moves.Size() != 8 {
		t.Errorf("Expected 8 moves, got %d", moves.Size())
	}
}
