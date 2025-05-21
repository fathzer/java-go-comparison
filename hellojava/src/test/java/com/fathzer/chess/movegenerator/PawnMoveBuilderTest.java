package com.fathzer.chess.movegenerator;

import org.junit.jupiter.api.Test;

class PawnMoveBuilderTest implements MoveBuilderTest {

    @Test
    void testWhitePawnInitialPosition() {
        testMoves("8/8/8/8/8/8/P7/8", "a2", "a3 a4", new PawnMoveBuilder(true));
    }

    @Test
    void testBlackPawnInitialPosition() {
        testMoves("1k6/p7/8/8/8/8/3K4/8", "a7", "a6 a5", new PawnMoveBuilder(false));
    }

    @Test
    void testWhitePawnWithCaptures() {
        testMoves("8/8/1p6/2P5/8/8/8/8", "c5", "c6 b6", new PawnMoveBuilder(true));
    }

    @Test
    void testBlackPawnWithCaptures() {
        testMoves("8/8/8/8/2p5/1P6/8/8", "c4", "c3 b3", new PawnMoveBuilder(false));
    }

    @Test
    void testBlockedPawn() {
        testMoves("8/8/8/8/8/2P5/2P5/8", "c2", "", new PawnMoveBuilder(true));
    }
}
