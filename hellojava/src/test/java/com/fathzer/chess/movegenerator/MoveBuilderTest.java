package com.fathzer.chess.movegenerator;

import static org.junit.jupiter.api.Assertions.*;

import java.util.LinkedList;
import java.util.List;

import com.fathzer.chess.Board;
import com.fathzer.chess.Move;

/**
 * Base class for testing MoveBuilder implementations.
 */
abstract class MoveBuilderTest {
    /**
     * Creates a MoveBuilder instance for testing.
     * @param isWhite true for white pieces, false for black
     * @return a new MoveBuilder instance
     */
    protected abstract MoveBuilder createBuilder(boolean isWhite);
    
    /**
     * Tests that the move builder generates the expected moves for a given position.
     * @param fen the FEN string representing the board position
     * @param fromSquare the square index (0-63) of the piece to move
     * @param isWhite true if the moving piece is white, false if black
     * @param expectedMoves the list of expected moves
     */
    protected void testMoves(String fen, int fromSquare, boolean isWhite, List<Move> expectedMoves) {
        // Create board and builder
        Board board = new Board(fen);
        System.out.println(board);
        MoveBuilder builder = createBuilder(isWhite);
        
        // Generate moves
        List<Move> moves = builder.build(new LinkedList<>(), board, fromSquare);
        
        // Verify moves
        assertEquals(expectedMoves.size(), moves.size(), 
            String.format("Expected %d moves but got %d: %s", 
                expectedMoves.size(), moves.size(), moves+" instead of "+expectedMoves));
                
        for (Move expected : expectedMoves) {
            assertTrue(moves.contains(expected), 
                String.format("Expected move %s not found in %s", expected, moves));
        }
    }
}
