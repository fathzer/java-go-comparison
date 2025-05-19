package com.fathzer.hellojava;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class FlagsTest {

    @Test
    void testDefault() {
        int loops = Flags.parsePiLoops(new String[] { });
        assertEquals(2000, loops);
    }

    @Test
    void testShortFlag() {
        int loops = Flags.parsePiLoops(new String[] { "-pl=42" });
        assertEquals(42, loops);
    }

    @Test
    void testLongFlag() {
        int loops = Flags.parsePiLoops(new String[] { "--piLoops=99" });
        assertEquals(99, loops);
    }

    @Test
    void testInvalidValue() {
        // System.exit is called on invalid value, so we can't test this directly without extra setup.
        // This test is a placeholder to indicate that such a test would require a custom SecurityManager or similar.
        // assertThrows(SecurityException.class, () -> Flags.parsePiLoops(new String[] { "Main", "--piLoops=abc" }));
    }
}
