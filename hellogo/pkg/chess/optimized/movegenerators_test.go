package optimized

import (
	"strings"
	"testing"
)

func TestKingInCorner(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/8/8/8/8/K7", "a1", "b1 a2 b2", Get(WhiteKing))
}

func TestKingInCenter(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/4K3/8/8/8/8", "e5", "d4 e4 f4 d5 f5 d6 e6 f6", Get(WhiteKing))
}

func TestKingWithBlockedSquares(t *testing.T) {
	// e4, f5 and f4 blocked by black pieces
	testMovesWithFEN(t, "K7/8/8/4kr2/3Ppn2/8/8/8", "e5", "d4 d5 d6 e6 f6", Get(BlackKing))
}

func TestRookInCorner(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/8/8/8/8/R7", "a1", "a2 a3 a4 a5 a6 a7 a8 b1 c1 d1 e1 f1 g1 h1", Get(WhiteRook))
}

func TestRookWithBlockedSquares(t *testing.T) {
	// Up (blocked by pawn at c3)
	testMovesWithFEN(t, "8/8/8/8/8/2P5/2R5/8", "c2", "b2 a2 d2 e2 f2 g2 h2 c1", Get(WhiteRook))
}

func TestBishopMoves(t *testing.T) {
	// Diagonal to h8 (stopped by black pawn at d4)
	testMovesWithFEN(t, "1k6/8/8/8/3p4/P7/1B6/K7", "b2", "c3 d4 c1", Get(WhiteBishop))
}

func TestKnightInCorner(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/8/8/8/8/N7", "a1", "c2 b3", Get(WhiteKnight))
}

func TestKnightInCenter(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/3N4/8/8/8/8", "d5", "b4 f4 c3 e3 f6 b6 c7 e7", Get(WhiteKnight))
}

func TestKnightWithBlockedSquares(t *testing.T) {
	testMovesWithFEN(t, "4k3/8/8/n1P5/3K4/1N6/3P4/8", "b3", "a1 a5 c1", Get(WhiteKnight))
}

func TestWhitePawnInitialPosition(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/8/8/8/P7/8", "a2", "a3 a4", Get(WhitePawn))
}

func TestBlackPawnInitialPosition(t *testing.T) {
	testMovesWithFEN(t, "1k6/p7/8/8/8/8/3K4/8", "a7", "a6 a5", Get(BlackPawn))
}

func TestWhitePawnWithCaptures(t *testing.T) {
	testMovesWithFEN(t, "8/8/1p6/2P5/8/8/8/8", "c5", "c6 b6", Get(WhitePawn))
}

func TestBlackPawnWithCaptures(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/8/2p5/1P6/8/8", "c4", "c3 b3", Get(BlackPawn))
}

func TestBlockedPawn(t *testing.T) {
	testMovesWithFEN(t, "8/8/8/8/8/2P5/2P5/8", "c2", "", Get(WhitePawn))
}

// testMovesWithFEN tests that the move builder generates the expected moves for a given position.
// fen: the FEN string representing the board position
// fromSquare: the square in UCI format (e.g., "e4")
// expectedDestinations: space-separated list of expected destination squares in UCI format
// builder: the move builder to test
func testMovesWithFEN(t *testing.T, fen, fromSquare, expectedDestinations string, builder MoveBuilder) {
	// Create board
	board := NewBoard(fen)

	// Generate moves
	moves := NewIntList()
	builder.Build(moves, board, GetSquare(fromSquare))

	// Parse expected moves
	expectedMoves := parseMoveList(board, fromSquare, expectedDestinations)

	// Verify moves
	testMoves(t, expectedMoves, moves)
}

// toString converts an IntList to a space-separated string of move notations
func intListToString(moves *IntList) string {
	var sb strings.Builder
	for i := 0; i < moves.Size(); i++ {
		if i > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(String(moves.Get(i)))
	}
	return sb.String()
}

// testMoves verifies that the generated moves match the expected moves
func testMoves(t *testing.T, expectedMoves, moves *IntList) {
	expectedSize := expectedMoves.Size()
	actualSize := moves.Size()
	if expectedSize != actualSize {
		t.Errorf("Expected %d moves but got %d: %s instead of %s",
			expectedSize, actualSize, intListToString(moves), intListToString(expectedMoves))
		return
	}

	for i := 0; i < expectedMoves.Size(); i++ {
		expectedMove := expectedMoves.Get(i)
		found := false
		for j := 0; j < moves.Size(); j++ {
			if moves.Get(j) == expectedMove {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected move %s not found in %s", String(expectedMove), intListToString(moves))
		}
	}
}

// parseMoveList parses a space-separated string of UCI moves into an IntList
func parseMoveList(board *Board, fromSquare, moveList string) *IntList {
	moves := NewIntList()
	if moveList == "" {
		return moves
	}

	from := GetSquare(fromSquare)
	for _, toSquare := range strings.Fields(moveList) {
		to := GetSquare(toSquare)
		capturedPiece := board.getPiece(to)
		moves.Add(Move(from, to, capturedPiece))
	}
	return moves
}
