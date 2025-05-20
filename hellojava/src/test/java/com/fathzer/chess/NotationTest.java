package com.fathzer.chess;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class NotationTest {
    @Test
    void testToUCI() {
        assertEquals("d1", Notation.toUCI(0, 3));
        assertEquals("e8", Notation.toUCI(7, 4));
    }
}
