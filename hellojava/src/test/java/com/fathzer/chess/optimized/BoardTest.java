package com.fathzer.chess.optimized;

import org.junit.jupiter.api.Test;

import static com.fathzer.chess.optimized.MoveGeneratorsTest.*;
import static org.junit.jupiter.api.Assertions.*;

import java.util.function.IntConsumer;
import java.util.stream.IntStream;

class BoardTest {

    @Test
    void testFenParsing() {
        // Standard chess starting position FEN (only piece placement part)
        String fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR";
        final Board board = new Board(fen);
        assertNotNull(board);
        assertEquals(Piece.WHITE_QUEEN, board.getPiece("d1"));

        // Test that blockers are set correctly
        IntConsumer expectBlocker = i -> assertEquals(Piece.BLOCKER, board.getPiece(i), "expected blocker at " + i);
        IntStream.range(0, 20).forEach(expectBlocker);
        IntStream.range(100, 120).forEach(expectBlocker);
    }

    @Test
    void testCopyConstructor() {
        String fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR";
        Board original = new Board(fen);
        Board copy = new Board(original);
        assertNotSame(original, copy);
    }

    @Test
    void testInvalidFenThrows() {
        String badFileCount = "rnbqkbnr/pppppppp/9/8/8/8/PPPPPPPP/RNBQKBNR"; // 9 is invalid
        assertThrows(IllegalArgumentException.class, () -> new Board(badFileCount));
        String badFileCount2 = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBN"; // Last row missed one file
        assertThrows(IllegalArgumentException.class, () -> new Board(badFileCount2));
        String badFileCount3 = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPP2/RNBQKBNR"; // 7th row has an extra file
        assertThrows(IllegalArgumentException.class, () -> new Board(badFileCount3));

        String badRankCount = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR/"; // Last rank is invalid
        assertThrows(IllegalArgumentException.class, () -> new Board(badRankCount));
        String badRankCount2 = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP"; // Last rank is missing
        assertThrows(IllegalArgumentException.class, () -> new Board(badRankCount2));

        String badPiece = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQXBNR"; // X is not a valid piece
        assertThrows(IllegalArgumentException.class, () -> new Board(badPiece));
    }
    
    @Test
    void testMakeMove() {
        // Start with initial position
        Board board = new Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");

        // Make a pawn move e2-e4
        int move = IntMove.fromUCI(board, "e2e4"); // e2-e4
        board.makeMove(move);
        
        // Verify the move was made
        assertEquals(Piece.NONE, board.getPiece("e2"), "Source square should be empty after move");
        assertEquals(Piece.WHITE_PAWN, board.getPiece("e4"), "Piece should be at destination");
    }
    
    @Test
    void testMakeCapture() {
        // Set up a position where white can capture a pawn
        Board board = new Board("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR");
        
        // Make a capture exd5
        int capture = IntMove.fromUCI(board, "e4d5"); // e4xd5
        board.makeMove(capture);
        
        // Verify the capture
        assertEquals(Piece.NONE, board.getPiece("e4"), "Source square should be empty after capture");
        assertEquals(Piece.WHITE_PAWN, board.getPiece("d5"), "Capturing piece should be at destination");
    }
    
    @Test
    void testUnmakeMove() {
        // Start with initial position
        Board original = new Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
        Board board = new Board(original);
        
        // Make a move
        int move = IntMove.fromUCI(board, "e2e4"); // e2-e4
        board.makeMove(move);
        
        // Unmake the move
        board.unmakeMove(move);
        
        // Board should be back to original state
        for (char rank = '1'; rank < '8'; rank++) {
            for (char file = 'a'; file < 'h'; file++) {
                assertEquals(original.getPiece(file + "" + rank), board.getPiece(file + "" + rank),
                    String.format("Pieces should match at %c%c", file, rank));
            }
        }
    }
    
    @Test
    void testUnmakeCapture() {
        // Set up a position where white can capture a pawn
        Board original = new Board("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR");
        Board board = new Board(original);
        
        // Make a capture
        int capture = IntMove.fromUCI(board, "e4d5"); // e4xd5
        board.makeMove(capture);
        
        // Unmake the capture
        board.unmakeMove(capture);
        
        // Board should be back to original state
        for (char rank = '1'; rank < '8'; rank++) {
            for (char file = 'a'; file < 'h'; file++) {
                assertEquals(original.getPiece(file + "" + rank), board.getPiece(file + "" + rank),
                    String.format("Pieces should match at %c%c", file, rank));
            }
        }
    }

    @Test
    void testGetMoves() {
        final Board board = new Board("8/8/8/8/1k6/8/pK6/Q7");
        final IntList moves = new IntList();
        board.getMoves(moves, true);
        IntList expected = new IntList();
        expected.addAll(parseMoveList(board, "a1", "a2 b1 c1 d1 e1 f1 g1 h1"));
        expected.addAll(parseMoveList(board, "b2", "a3 b3 c3 a2 c2 b1 c1"));
        testMoves(expected, moves);
        board.getMoves(moves, false);
        testMoves(parseMoveList(board, "b4", "a5 b5 c5 a4 c4 a3 b3 c3"), moves);
        assertEquals(8, moves.size());
    }
}
