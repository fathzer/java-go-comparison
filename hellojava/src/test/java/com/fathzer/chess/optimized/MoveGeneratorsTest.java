package com.fathzer.chess.optimized;

import static com.fathzer.chess.optimized.MoveGenerators.get;
import static com.fathzer.chess.optimized.Piece.*;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

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
        final IntList moves = new IntList();
        builder.build(moves, board, Board.getSquare(fromSquare));
        
        // Verify moves
        testMoves(parseMoveList(board, fromSquare, expectedDestinations), moves);
    }

    static void testMoves(IntList expectedMoves, IntList moves) {
        assertEquals(expectedMoves.size(), moves.size(), 
            String.format("Expected %d moves but got %d: %s", 
                expectedMoves.size(), moves.size(), moves+" instead of "+toString(expectedMoves)));
        
        for (int i = 0; i < expectedMoves.size(); i++) {
            assertTrue(moves.contains(expectedMoves.get(i)), 
                String.format("Expected move %s not found in %s", IntMove.toString(expectedMoves.get(i)), toString(moves)));
        }
    }
    
    /**
     * Converts an IntList to a space-separated string of move notations.
     * 
     * @param moves the moves to convert
     * @return a string representation of the moves
     */
    private static String toString(IntList moves) {
        return IntStream.range(0, moves.size()).map(moves::get).mapToObj(IntMove::toString).collect(Collectors.joining(" "));
    }

    static IntList parseMoveList(Board board, String fromSquare, String moveList) {
        final IntList moves = new IntList();
        if (moveList.isEmpty()) {
            return moves;
        }
        final int from = Board.getSquare(fromSquare);
        Arrays.stream(moveList.split(" ")).mapToInt(Board::getSquare).map(toSquare -> {
            final int capturedPiece = board.getPiece(toSquare);
            return IntMove.move(from, toSquare, capturedPiece);
        }).forEach(moves::add);
        return moves;
    }

}
