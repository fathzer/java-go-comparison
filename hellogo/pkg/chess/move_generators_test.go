package chess

import (
	"strings"
	"testing"
)

func TestKingInCorner(t *testing.T) {
	testMoves(t, "8/8/8/8/8/8/8/K7", "a1", "b1 a2 b2", GetMoveGenerator(&WHITE_KING))
}

func TestKingInCenter(t *testing.T) {
	testMoves(t, "8/8/8/4K3/8/8/8/8", "e5", "d4 e4 f4 d5 f5 d6 e6 f6", GetMoveGenerator(&WHITE_KING))
}

func TestKingWithBlockedSquares(t *testing.T) {
	// e4, f5 and f4 blocked by black pieces
	testMoves(t, "K7/8/8/4kr2/3Ppn2/8/8/8", "e5", "d4 d5 d6 e6 f6", GetMoveGenerator(&BLACK_KING))
}

func TestRookInCorner(t *testing.T) {
	expected := "a2 a3 a4 a5 a6 a7 a8 b1 c1 d1 e1 f1 g1 h1"
	testMoves(t, "8/8/8/8/8/8/8/R7", "a1", expected, GetMoveGenerator(&WHITE_ROOK))
}

func TestRookWithBlockedSquares(t *testing.T) {
	// Up (blocked by pawn at c3)
	testMoves(t, "8/8/8/8/8/2P5/2R5/8", "c2", "b2 a2 d2 e2 f2 g2 h2 c1", GetMoveGenerator(&WHITE_ROOK))
}

func TestBishopMoves(t *testing.T) {
	// Diagonal to h8 (stopped by black pawn at d4)
	testMoves(t, "1k6/8/8/8/3p4/P7/1B6/K7", "b2", "c3 d4 c1", GetMoveGenerator(&WHITE_BISHOP))
}

func TestKnightInCorner(t *testing.T) {
	testMoves(t, "8/8/8/8/8/8/8/N7", "a1", "c2 b3", GetMoveGenerator(&WHITE_KNIGHT))
}

func TestKnightInCenter(t *testing.T) {
	testMoves(t, "8/8/8/3N4/8/8/8/8", "d5", "b4 f4 c3 e3 f6 b6 c7 e7", GetMoveGenerator(&WHITE_KNIGHT))
}

func TestKnightWithBlockedSquares(t *testing.T) {
	testMoves(t, "4k3/8/8/n1P5/3K4/1N6/3P4/8", "b3", "a1 a5 c1", GetMoveGenerator(&WHITE_KNIGHT))
}

func TestWhitePawnInitialPosition(t *testing.T) {
	testMoves(t, "8/8/8/8/8/8/P7/8", "a2", "a3 a4", GetMoveGenerator(&WHITE_PAWN))
}

func TestBlackPawnInitialPosition(t *testing.T) {
	testMoves(t, "1k6/p7/8/8/8/8/3K4/8", "a7", "a6 a5", GetMoveGenerator(&BLACK_PAWN))
}

func TestWhitePawnWithCaptures(t *testing.T) {
	testMoves(t, "8/8/1p6/2P5/8/8/8/8", "c5", "c6 b6", GetMoveGenerator(&WHITE_PAWN))
}

func TestBlackPawnWithCaptures(t *testing.T) {
	testMoves(t, "8/8/8/8/2p5/1P6/8/8", "c4", "c3 b3", GetMoveGenerator(&BLACK_PAWN))
}

func TestBlockedPawn(t *testing.T) {
	testMoves(t, "8/8/8/8/8/2P5/2P5/8", "c2", "", GetMoveGenerator(&WHITE_PAWN))
}

// testMoves tests that the move builder generates the expected moves for a given position.
func testMoves(t *testing.T, fen, fromSquare, expectedDestinations string, builder MoveBuilder) {
	t.Helper()

	// Create board
	board, err := NewBoard(fen)
	if err != nil {
		t.Fatalf("Failed to create board: %v", err)
	}

	// Generate moves
	moves := make([]Move, 0, 30)
	builder.Build(&moves, board, GetSquare(fromSquare))

	// Parse expected moves
	expectedMoves := parseMoveList(t, fromSquare, expectedDestinations)

	// Verify moves
	t.Run(fen, func(t *testing.T) {
		if len(expectedMoves) != len(moves) {
			t.Errorf("Expected %d moves but got %d: %v instead of %v", 
				len(expectedMoves), len(moves), movesToString(moves), expectedDestinations)
			return
		}

		for _, expected := range expectedMoves {
			found := false
			for _, m := range moves {
				if m == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected move %v not found in %v", expected, movesToString(moves))
			}
		}

		// Also check for unexpected moves
		for _, m := range moves {
			found := false
			for _, expected := range expectedMoves {
				if m == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Unexpected move %v found in %v", m, movesToString(moves))
			}
		}
	})
}

// parseMoveList parses a space-separated list of destination squares into Move objects.
func parseMoveList(t *testing.T, fromSquare, moveList string) []Move {
	t.Helper()
	if moveList == "" {
		return nil
	}
	squares := strings.Fields(moveList)
	moves := make([]Move, 0, len(squares))
	for _, toSquare := range squares {
		moves = append(moves, NewMove(GetSquare(fromSquare), GetSquare(toSquare)))
	}
	return moves
}

// movesToString converts a slice of moves to a space-separated string of destination squares.
func movesToString(moves []Move) string {
	if len(moves) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(GetUCI(moves[0].To()))
	for i := 1; i < len(moves); i++ {
		sb.WriteByte(' ')
		sb.WriteString(GetUCI(moves[i].To()))
	}
	return sb.String()
}
