package com.fathzer.pi;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class PiCalculatorTest {

    @Test
    void testComputePiWithValidDigits() {
        String pi = PiCalculator.computePi(10).toString();
        assertNotNull(pi);
        assertTrue(pi.startsWith("3.14159"), "Pi should start with 3.14159");
        assertEquals(10, pi.substring(2).length(), "Should have 10 digits after decimal point");
    }

    @Test
    void testComputePiAccuracy() {
        String pi = PiCalculator.computePi(3).toString();
        // Pi to 3 decimal digits is 3.141
        assertEquals("3.141", pi);
    }

    @Test
    void testComputePiThrowsOnZeroOrNegative() {
        assertThrows(IllegalArgumentException.class, () -> PiCalculator.computePi(0));
        assertThrows(IllegalArgumentException.class, () -> PiCalculator.computePi(-10));
    }
}
