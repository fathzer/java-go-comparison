package com.fathzer.chess.movegenerator;

import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

import com.fathzer.chess.Move;

class KingMoveBuilderTest extends MoveBuilderTest {
    @Override
    protected MoveBuilder createBuilder(boolean isWhite) {
        return new KingMoveBuilder(isWhite);
    }

    @Test
    void testKingInCorner() {
        String fen = "8/8/8/8/8/8/8/K7";
        int fromSquare = 0; // a1
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(0, 1),  // a1-b1
            new Move(0, 8),  // a1-a2
            new Move(0, 9)   // a1-b2
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testKingInCenter() {
        String fen = "8/8/8/4K3/8/8/8/8";
        int fromSquare = 36; // e5
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(36, 27), // e5-d4
            new Move(36, 28), // e5-e4
            new Move(36, 29), // e5-f4
            new Move(36, 35), // e5-d5
            new Move(36, 37), // e5-f5
            new Move(36, 43), // e5-d6
            new Move(36, 44), // e5-e6
            new Move(36, 45)  // e5-f6
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testKingWithBlockedSquares() {
        String fen = "8/8/8/4k3/3P4/4P3/8/8";
        int fromSquare = 28; // e4
        boolean isWhite = true; // The king at e4 is white
        List<Move> expected = Arrays.asList(
            // e4-d3 (blocked by white pawn)
            // e4-e3 (blocked by white pawn)
            // e4-f3 (blocked by black king)
            new Move(28, 19), // e4-d5
            new Move(28, 20), // e4-e5
            new Move(28, 21)  // e4-f5
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }
}
