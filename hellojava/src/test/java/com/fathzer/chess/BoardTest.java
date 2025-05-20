package com.fathzer.chess;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class BoardTest {

    @Test
    void testFenParsing() {
        // Standard chess starting position FEN (only piece placement part)
        String fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR";
        Board board = new Board(fen);
        assertNotNull(board);
        assertEquals(Piece.WHITE_QUEEN, board.getPiece(board.getSquare(0, 3)));
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
        Move move = new Move(1, 4, 3, 4); // e2-e4
        board.makeMove(move);
        
        // Verify the move was made
        assertNull(board.getPiece(board.getSquare(1, 4)), "Source square should be empty after move");
        assertEquals(Piece.WHITE_PAWN, board.getPiece(board.getSquare(3, 4)), "Piece should be at destination");
    }
    
    @Test
    void testMakeCapture() {
        // Set up a position where white can capture a pawn
        Board board = new Board("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR");
        
        // Make a capture exd5
        Move capture = new Move(3, 4, 4, 3); // e4xd5
        board.makeMove(capture);
        
        // Verify the capture
        assertNull(board.getPiece(board.getSquare(3, 4)), "Source square should be empty after capture");
        assertEquals(Piece.WHITE_PAWN, board.getPiece(board.getSquare(4, 3)), "Capturing piece should be at destination");
    }
    
    @Test
    void testUnmakeMove() {
        // Start with initial position
        Board original = new Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
        Board board = new Board(original);
        
        // Make a move
        Move move = new Move(6, 4, 4, 4); // e2-e4
        board.makeMove(move);
        
        // Unmake the move
        board.unmakeMove();
        
        // Board should be back to original state
        for (int rank = 0; rank < 8; rank++) {
            for (int file = 0; file < 8; file++) {
                assertEquals(original.getPiece(board.getSquare(rank, file)), board.getPiece(board.getSquare(rank, file)),
                    String.format("Pieces should match at %c%c", 'a' + file, '1' + rank));
            }
        }
    }
    
    @Test
    void testUnmakeCapture() {
        // Set up a position where white can capture a pawn
        Board original = new Board("rnbqkbnr/ppp1pppp/8/3p4/4P3/8/PPPP1PPP/RNBQKBNR");
        Board board = new Board(original);
        
        // Make a capture
        Move capture = new Move(4, 4, 3, 3); // e4xd5
        board.makeMove(capture);
        
        // Unmake the capture
        board.unmakeMove();
        
        // Board should be back to original state
        for (int rank = 0; rank < 8; rank++) {
            for (int file = 0; file < 8; file++) {
                assertEquals(original.getPiece(board.getSquare(rank, file)), board.getPiece(board.getSquare(rank, file)),
                    String.format("Pieces should match at %c%c", 'a' + file, '1' + rank));
            }
        }
    }
    
    @Test
    void testUnmakeWithoutMoves() {
        Board board = new Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
        assertThrows(IllegalStateException.class, board::unmakeMove,
            "Should throw when no moves to unmake"
        );
    }
}
