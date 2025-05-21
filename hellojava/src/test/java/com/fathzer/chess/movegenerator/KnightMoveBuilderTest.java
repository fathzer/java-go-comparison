package com.fathzer.chess.movegenerator;

import org.junit.jupiter.api.Test;

class KnightMoveBuilderTest implements MoveBuilderTest {
    private static final MoveBuilder BUILDER = new KnightMoveBuilder(true);

    @Test
    void testKnightInCorner() {
        testMoves("8/8/8/8/8/8/8/N7", "a1", "c2 b3", BUILDER);
    }

    @Test
    void testKnightInCenter() {
        testMoves("8/8/8/3N4/8/8/8/8", "d5", "b4 f4 c3 e3 f6 b6 c7 e7", BUILDER);
    }

    @Test
    void testKnightWithBlockedSquares() {
        testMoves("4k3/8/8/n1P5/3K4/1N6/3P4/8", "b3", "a1 a5 c1", BUILDER);
    }
}
