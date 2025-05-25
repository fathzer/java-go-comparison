package com.fathzer.chess.optimized;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

import static org.junit.jupiter.api.Assertions.*;

class IntMoveTest {

    @ParameterizedTest
    @CsvSource({
        // from, to, capture
        "0, 0, 0",
        "127, 127, 31",
        "64, 32, 15",
        "1, 2, 1",
        "100, 50, 0",
        "10, 20, 31"
    })
    void testEncodeDecode(int from, int to, int capture) {
        // Test move creation with capture
        int move = IntMove.move(from, to, capture);
        assertEquals(from, IntMove.from(move), "from mismatch");
        assertEquals(to, IntMove.to(move), "to mismatch");
        assertEquals(capture, IntMove.capture(move), "capture mismatch");
        
        // Test that other bits don't affect extraction
        int modifiedMove = move | 0x80000000;
        assertEquals(capture, IntMove.capture(modifiedMove), "capture mismatch with high bit set");
    }

    @Test
    void testCapture() {
        // Test capture extraction
        int move = IntMove.move(10, 20, 15);
        assertEquals(15, IntMove.capture(move));
        
        // Test that other bits don't affect capture value
        int modifiedMove = move | 0x80000000; // Set high bit
        assertEquals(15, IntMove.capture(modifiedMove));
        
        // Test default capture is 0
        int defaultMove = IntMove.move(10, 20, 0);
        assertEquals(0, IntMove.capture(defaultMove));
    }
    
    @ParameterizedTest
    @CsvSource({
        // from, to, captured piece, expectedUCI
        "21, 31, 0, 'a1a2'",
        "31, 21, 1, 'a2xa1(P)'",
        "88, 87, 0, 'h7g7'",
        "98, 97, 11, 'h8xg8(q)'",
        "21, 98, 0, 'a1h8'",
        "98, 21, 0, 'h8a1'"
    })
    void testToString(int from, int to, int capture, String expectedUCI) {
        int move = IntMove.move(from, to, capture);
        assertEquals(expectedUCI, IntMove.toString(move), "UCI string mismatch");
    }
}
