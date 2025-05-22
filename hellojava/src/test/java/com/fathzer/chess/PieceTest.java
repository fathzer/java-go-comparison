package com.fathzer.chess;

import static org.junit.jupiter.api.Assertions.*;

import java.util.stream.Stream;

import static com.fathzer.chess.Piece.*;

import org.junit.jupiter.api.Test;

class PieceTest {

    @Test
    void testFromCode() {
        // Test all valid piece codes
        assertEquals(WHITE_PAWN, fromCode('P'));
        assertEquals(WHITE_KNIGHT, fromCode('N'));
        assertEquals(WHITE_BISHOP, fromCode('B'));
        assertEquals(WHITE_ROOK, fromCode('R'));
        assertEquals(WHITE_QUEEN, fromCode('Q'));
        assertEquals(WHITE_KING, fromCode('K'));
        assertEquals(BLACK_PAWN, fromCode('p'));
        assertEquals(BLACK_KNIGHT, fromCode('n'));
        assertEquals(BLACK_BISHOP, fromCode('b'));
        assertEquals(BLACK_ROOK, fromCode('r'));
        assertEquals(BLACK_QUEEN, fromCode('q'));
        assertEquals(BLACK_KING, fromCode('k'));
        
        // Test invalid code
        assertNull(fromCode('X'));
        assertNull(fromCode(' '));
    }

    @Test
    void testGetCode() {
        // Test a sample of pieces
        assertEquals('P', WHITE_PAWN.getCode());
        assertEquals('N', WHITE_KNIGHT.getCode());
        assertEquals('p', BLACK_PAWN.getCode());
        assertEquals('k', BLACK_KING.getCode());
    }

    @Test
    void testIsWhite() {
        // Test white pieces
        Stream.of(WHITE_PAWN, WHITE_KNIGHT, WHITE_BISHOP, WHITE_ROOK, WHITE_QUEEN, WHITE_KING)
            .forEach(piece -> assertTrue(piece.isWhite()));
        
        // Test black pieces
        Stream.of(BLACK_PAWN, BLACK_KNIGHT, BLACK_BISHOP, BLACK_ROOK, BLACK_QUEEN, BLACK_KING)
            .forEach(piece -> assertFalse(piece.isWhite()));
    }

    @Test
    void testCanBeCapturedBy() {
        // Test white pieces can be captured by black and not by white
        Stream.of(WHITE_PAWN, WHITE_KNIGHT, WHITE_BISHOP, WHITE_ROOK, WHITE_QUEEN, WHITE_KING)
            .forEach(piece -> {
                assertTrue(piece.canBeCapturedBy(false));
                assertFalse(piece.canBeCapturedBy(true));
            });
        // Test black pieces can be captured by white and not by black
        Stream.of(BLACK_PAWN, BLACK_KNIGHT, BLACK_BISHOP, BLACK_ROOK, BLACK_QUEEN, BLACK_KING)
            .forEach(piece -> {
                assertTrue(piece.canBeCapturedBy(true));
                assertFalse(piece.canBeCapturedBy(false));
            });        
   }

    @Test
    void testBlocker() {
        // Test BLOCKER specific properties
        assertFalse(BLOCKER.canBeCapturedBy(true));
        assertFalse(BLOCKER.canBeCapturedBy(false));
    }
}
