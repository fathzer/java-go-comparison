package com.fathzer.chess.movegenerator;

import static org.junit.jupiter.api.Assertions.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.junit.jupiter.api.Test;

import com.fathzer.chess.Board;
import com.fathzer.chess.Direction;
import com.fathzer.chess.Move;

class SliderMoveBuilderTest extends MoveBuilderTest {
    @Override
    protected MoveBuilder createBuilder(boolean isWhite) {
        // Test with rook moves (orthogonal)
        return new SliderMoveBuilder(
            new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST },
            isWhite
        );
    }

    @Test
    void testRookInCorner() {
        String fen = "8/8/8/8/8/8/8/R7";
        int fromSquare = 0; // a1
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            // Horizontal
            new Move(0, 1), new Move(0, 2), new Move(0, 3), new Move(0, 4), 
            new Move(0, 5), new Move(0, 6), new Move(0, 7),
            // Vertical
            new Move(0, 8), new Move(0, 16), new Move(0, 24), new Move(0, 32), 
            new Move(0, 40), new Move(0, 48), new Move(0, 56)
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testRookWithBlockedSquares() {
        String fen = "8/8/8/8/8/2P5/2R5/8";
        int fromSquare = 10; // c2
        boolean isWhite = true;
        List<Move> expected = Arrays.asList(
            // Left
            new Move(10, 9), // c2-b2
            // Right
            new Move(10, 11), // c2-d2
            new Move(10, 12), // c2-e2
            new Move(10, 13), // c2-f2
            new Move(10, 14), // c2-g2
            new Move(10, 15), // c2-h2
            // Up (blocked by pawn at c3)
            // Down
            new Move(10, 2),  // c2-c1
            new Move(10, 18), // c2-c3 (blocked by white pawn)
            new Move(10, 26), // c2-c4
            new Move(10, 34), // c2-c5
            new Move(10, 42), // c2-c6
            new Move(10, 50), // c2-c7
            new Move(10, 58)  // c2-c8
        );
        testMoves(fen, fromSquare, isWhite, expected);
    }

    @Test
    void testBishopMoves() {
        // Test bishop moves (diagonal)
        SliderMoveBuilder bishopBuilder = new SliderMoveBuilder(
            new Direction[] { 
                Direction.NORTH_EAST, Direction.NORTH_WEST, 
                Direction.SOUTH_EAST, Direction.SOUTH_WEST 
            },
            true
        );
        
        String fen = "8/8/8/3p4/4P3/8/8/B7";
        Board board = new Board(fen);
        int fromSquare = 0; // a1
        
        List<Move> moves = bishopBuilder.build(new ArrayList<>(), board, fromSquare);
        
        List<Move> expected = Arrays.asList(
            // Diagonal to h8 (blocked by black pawn at d5)
            new Move(0, 9),  // a1-b2
            new Move(0, 18), // a1-c3
            new Move(0, 27), // a1-d4
            // Diagonal to h1
            new Move(0, 9),  // a1-b2
            new Move(0, 18), // a1-c3
            new Move(0, 27), // a1-d4
            // Diagonal to a8 (blocked by white pawn at e4)
            new Move(0, 9),  // a1-b2
            new Move(0, 18)  // a1-c3
        );
        
        assertEquals(expected.size(), moves.size(), 
            String.format("Expected %d moves but got %d: %s", 
                expected.size(), moves.size(), moves));
                
        for (Move move : expected) {
            assertTrue(moves.contains(move), 
                String.format("Expected move %s not found in %s", move, moves));
        }
    }
}
