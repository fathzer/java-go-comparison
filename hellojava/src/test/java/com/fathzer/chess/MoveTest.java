package com.fathzer.chess;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

class MoveTest {
    @Test
    void testValidMove() {
        // Test valid move construction
        Move move = new Move(0, 16); // a1 to a3
        assertEquals(0, move.from());
        assertEquals(16, move.to());
    }

    
    @Test
    void testInvalidMove() {
        // Test invalid from/to indices
        assertThrows(IllegalArgumentException.class, () -> new Move(-1, 0));
        assertThrows(IllegalArgumentException.class, () -> new Move(0, -1));
        assertThrows(IllegalArgumentException.class, () -> new Move(64, 0));
        assertThrows(IllegalArgumentException.class, () -> new Move(0, 64));
        assertThrows(IllegalArgumentException.class, () -> new Move(10, 10));
    }
    
    @Test
    void testToString() {
        // Test standard move
        Move e2e4 = new Move(1, 4 , 3, 4); // e2 to e4
        assertEquals("e2e4", e2e4.toString());
    }
    
    @Test
    void testEquality() {
        // Test record's built-in equality
        Move move1 = new Move(0, 16);
        Move move2 = new Move(0, 16);
        Move move3 = new Move(16, 0);
        
        assertEquals(move1, move2);
        assertEquals(move1.hashCode(), move2.hashCode());
        assertNotEquals(move1, move3);
    }
}
