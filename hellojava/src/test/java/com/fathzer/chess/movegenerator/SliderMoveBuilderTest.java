package com.fathzer.chess.movegenerator;

import org.junit.jupiter.api.Test;

import com.fathzer.chess.Piece;

class SliderMoveBuilderTest implements MoveBuilderTest {
    @Test
    void testRookInCorner() {
        testMoves("8/8/8/8/8/8/8/R7", "a1", "a2 a3 a4 a5 a6 a7 a8 b1 c1 d1 e1 f1 g1 h1", Piece.WHITE_ROOK.getMoveBuilder());
    }

    @Test
    void testRookWithBlockedSquares() {
        // Up (blocked by pawn at c3)
        testMoves("8/8/8/8/8/2P5/2R5/8", "c2", "b2 a2 d2 e2 f2 g2 h2 c1", Piece.WHITE_ROOK.getMoveBuilder());
    }

    @Test
    void testBishopMoves() {
        // Diagonal to h8 (stopped by black pawn at d4)
        testMoves("1k6/8/8/8/3p4/P7/1B6/K7", "b2", "c3 d4 c1", Piece.WHITE_BISHOP.getMoveBuilder());
    }
}
