package com.fathzer.chess.movegenerator;

import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

import com.fathzer.chess.Move;

class KnightMoveBuilderTest extends MoveBuilderTest {
    @Override
    protected MoveBuilder createBuilder(boolean isWhite) {
        return new KnightMoveBuilder(isWhite);
    }

    @Test
    void testKnightInCorner() {
        String fen = "8/8/8/8/8/8/8/N7";
        int fromSquare = 0; // a1
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(0, 10), // a1-c2
            new Move(0, 17)  // a1-b3
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testKnightInCenter() {
        String fen = "8/8/8/3N4/8/8/8/8";
        int fromSquare = 35; // d5
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(35, 18), // d5-b4
            new Move(35, 20), // d5-f4
            new Move(35, 25), // d5-c3
            new Move(35, 29), // d5-e3
            new Move(35, 41), // d5-f6
            new Move(35, 45), // d5-b6
            new Move(35, 50), // d5-c7
            new Move(35, 52)  // d5-e7
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testKnightWithBlockedSquares() {
        String fen = "8/8/8/3n4/2P1P3/1N6/8/8";
        int fromSquare = 17; // b3
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(17, 2),  // b3-a1
            new Move(17, 32), // b3-c1
            new Move(17, 3),  // b3-c1 (should be blocked by pawn)
            new Move(17, 20), // b3-d2
            new Move(17, 27), // b3-d4 (blocked by black knight)
            new Move(17, 33), // b3-a5
            new Move(17, 34)  // b3-c5 (blocked by pawn)
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }
}
