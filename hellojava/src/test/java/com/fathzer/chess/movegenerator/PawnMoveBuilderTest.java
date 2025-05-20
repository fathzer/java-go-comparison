package com.fathzer.chess.movegenerator;

import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

import com.fathzer.chess.Move;

class PawnMoveBuilderTest extends MoveBuilderTest {
    @Override
    protected MoveBuilder createBuilder(boolean isWhite) {
        return new PawnMoveBuilder(isWhite);
    }

    @Test
    void testWhitePawnInitialPosition() {
        String fen = "8/8/8/8/8/8/P7/8";
        int fromSquare = 8; // a2
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(8, 16),  // a2-a3
            new Move(8, 24)   // a2-a4 (double move)
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testBlackPawnInitialPosition() {
        String fen = "8/p7/8/8/8/8/8/8";
        int fromSquare = 48; // a7
        boolean isWhite = false;
        List<Move> expected = Arrays.asList(
            new Move(48, 40),  // a7-a6
            new Move(48, 32)   // a7-a5 (double move)
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testWhitePawnWithCaptures() {
        String fen = "8/8/1p6/2P5/8/8/8/8";
        int fromSquare = 34; // c5
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            new Move(34, 42),  // c5-c6
            new Move(34, 41)   // c5xb6
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testBlackPawnWithCaptures() {
        String fen = "8/8/8/8/2p5/1P6/8/8";
        int fromSquare = 26; // c4
        boolean isWhite = false;
        List<Move> expected = Arrays.asList(
            new Move(26, 18),  // c4-c3
            new Move(26, 17)   // c4xb3
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testBlockedPawn() {
        String fen = "8/8/8/8/8/2P5/2P5/8";
        int fromSquare = 10; // c2
        boolean isWhite = true;
        List<Move> expected = List.of();  // No moves possible
        testMoves(fen, fromSquare, isWhite, expected);
    }
}
