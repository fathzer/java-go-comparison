package com.fathzer.chess.movegenerator;

import org.junit.jupiter.api.Test;

class KingMoveBuilderTest implements MoveBuilderTest {
    private static final MoveBuilder BUILDER = new KingMoveBuilder(true);

    @Test
    void testKingInCorner() {
        testMoves("8/8/8/8/8/8/8/K7", "a1", "b1 a2 b2", BUILDER);
    }

    @Test
    void testKingInCenter() {
        testMoves("8/8/8/4K3/8/8/8/8", "e5", "d4 e4 f4 d5 f5 d6 e6 f6", BUILDER);
    }

    @Test
    void testKingWithBlockedSquares() {
        // e4, f5 and f4 blocked by black pieces
        testMoves("K7/8/8/4kr2/3Ppn2/8/8/8", "e5", "d4 d5 d6 e6 f6", new KingMoveBuilder(false));
    }
}
