package com.fathzer.chess.movegenerator;

import static org.junit.jupiter.api.Assertions.*;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

import com.fathzer.chess.Board;
import com.fathzer.chess.Move;

/**
 * Base class for testing MoveBuilder implementations.
 */
public interface MoveBuilderTest {
    
    /**
     * Tests that the move builder generates the expected moves for a given position.
     * @param fen the FEN string representing the board position
     * @param fromSquare the square index (0-63) of the piece to move
     * @param expectedDestinations the list of expected destinations in UCI format separated by spaces
     * @param builder the move builder to test
     */
    default void testMoves(String fen, String fromSquare, String expectedDestinations, MoveBuilder builder) {
        // Create board and builder
        final Board board = new Board(fen);
        
        // Generate moves
        final List<Move> moves = builder.build(new LinkedList<>(), board, Board.getSquare(fromSquare));
        
        // Verify moves
        testMoves(parseMoveList(fromSquare, expectedDestinations), moves);
    }

    default void testMoves(List<Move> expectedMoves, List<Move> moves) {
        assertEquals(expectedMoves.size(), moves.size(), 
            String.format("Expected %d moves but got %d: %s", 
                expectedMoves.size(), moves.size(), moves+" instead of "+expectedMoves));
        
        for (Move expected : expectedMoves) {
            assertTrue(moves.contains(expected), 
                String.format("Expected move %s not found in %s", expected, moves));
        }
    }

    default List<Move> parseMoveList(String fromSquare, String moveList) {
        if (moveList.isEmpty()) {
            return List.of();
        }
        return Arrays.stream(moveList.split(" ")).map(toSquare -> new Move(fromSquare, toSquare)).toList();
    }
}
