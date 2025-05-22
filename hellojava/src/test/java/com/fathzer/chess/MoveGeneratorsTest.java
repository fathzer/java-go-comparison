package com.fathzer.chess;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

import static com.fathzer.chess.Piece.*;
import static com.fathzer.chess.MoveGenerators.get;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

class MoveGeneratorsTest {
    @Test
    void testKingInCorner() {
        testMoves("8/8/8/8/8/8/8/K7", "a1", "b1 a2 b2", get(WHITE_KING));
    }

    @Test
    void testKingInCenter() {
        testMoves("8/8/8/4K3/8/8/8/8", "e5", "d4 e4 f4 d5 f5 d6 e6 f6", get(WHITE_KING));
    }

    @Test
    void testKingWithBlockedSquares() {
        // e4, f5 and f4 blocked by black pieces
        testMoves("K7/8/8/4kr2/3Ppn2/8/8/8", "e5", "d4 d5 d6 e6 f6", get(BLACK_KING));
    }

    @Test
    void testRookInCorner() {
        testMoves("8/8/8/8/8/8/8/R7", "a1", "a2 a3 a4 a5 a6 a7 a8 b1 c1 d1 e1 f1 g1 h1", get(WHITE_ROOK));
    }

    @Test
    void testRookWithBlockedSquares() {
        // Up (blocked by pawn at c3)
        testMoves("8/8/8/8/8/2P5/2R5/8", "c2", "b2 a2 d2 e2 f2 g2 h2 c1", get(WHITE_ROOK));
    }

    @Test
    void testBishopMoves() {
        // Diagonal to h8 (stopped by black pawn at d4)
        testMoves("1k6/8/8/8/3p4/P7/1B6/K7", "b2", "c3 d4 c1", get(WHITE_BISHOP));
    }

    @Test
    void testKnightInCorner() {
        testMoves("8/8/8/8/8/8/8/N7", "a1", "c2 b3", get(WHITE_KNIGHT));
    }

    @Test
    void testKnightInCenter() {
        testMoves("8/8/8/3N4/8/8/8/8", "d5", "b4 f4 c3 e3 f6 b6 c7 e7", get(WHITE_KNIGHT));
    }

    @Test
    void testKnightWithBlockedSquares() {
        testMoves("4k3/8/8/n1P5/3K4/1N6/3P4/8", "b3", "a1 a5 c1", get(WHITE_KNIGHT));
    }
    
    @Test
    void testWhitePawnInitialPosition() {
        testMoves("8/8/8/8/8/8/P7/8", "a2", "a3 a4", get(WHITE_PAWN));
    }

    @Test
    void testBlackPawnInitialPosition() {
        testMoves("1k6/p7/8/8/8/8/3K4/8", "a7", "a6 a5", get(BLACK_PAWN));
    }

    @Test
    void testWhitePawnWithCaptures() {
        testMoves("8/8/1p6/2P5/8/8/8/8", "c5", "c6 b6", get(WHITE_PAWN));
    }

    @Test
    void testBlackPawnWithCaptures() {
        testMoves("8/8/8/8/2p5/1P6/8/8", "c4", "c3 b3", get(BLACK_PAWN));
    }

    @Test
    void testBlockedPawn() {
        testMoves("8/8/8/8/8/2P5/2P5/8", "c2", "", get(WHITE_PAWN));
    }

    /**
     * Tests that the move builder generates the expected moves for a given position.
     * @param fen the FEN string representing the board position
     * @param fromSquare the square index (0-63) of the piece to move
     * @param expectedDestinations the list of expected destinations in UCI format separated by spaces
     * @param builder the move builder to test
     */
    static void testMoves(String fen, String fromSquare, String expectedDestinations, MoveBuilder builder) {
        // Create board and builder
        final Board board = new Board(fen);
        
        // Generate moves
        final List<Move> moves = new LinkedList<>();
        builder.build(moves, board.explorable, Board.getSquare(fromSquare), Move::new);
        
        // Verify moves
        testMoves(parseMoveList(fromSquare, expectedDestinations), moves);
    }

    static void testMoves(List<Move> expectedMoves, List<Move> moves) {
        assertEquals(expectedMoves.size(), moves.size(), 
            String.format("Expected %d moves but got %d: %s", 
                expectedMoves.size(), moves.size(), moves+" instead of "+expectedMoves));
        
        for (Move expected : expectedMoves) {
            assertTrue(moves.contains(expected), 
                String.format("Expected move %s not found in %s", expected, moves));
        }
    }

    static List<Move> parseMoveList(String fromSquare, String moveList) {
        if (moveList.isEmpty()) {
            return List.of();
        }
        return Arrays.stream(moveList.split(" ")).map(toSquare -> new Move(fromSquare, toSquare)).toList();
    }

}
