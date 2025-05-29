package com.fathzer.chess.common;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

class PerftResultTest {

    @Test
    void testDefaultConstructor() {
        PerftResult result = new PerftResult();
        assertEquals(0, result.leafNodesCount(), "Default leafNodesCount should be 0");
        assertEquals(0, result.searchedNodesCount(), "Default searchedNodesCount should be 0");
    }

    @Test
    void testSettersAndGetters() {
        PerftResult result = new PerftResult();
        
        // Test setLeafNodesCount and leafNodesCount
        result.setLeafNodesCount(42);
        assertEquals(42, result.leafNodesCount(), "leafNodesCount should be 42");
        
        // Test incrementSearchedNodesCount and searchedNodesCount
        result.incrementSearchedNodesCount();
        assertEquals(1, result.searchedNodesCount(), "searchedNodesCount should be 1 after increment");
        
        result.incrementSearchedNodesCount();
        assertEquals(2, result.searchedNodesCount(), "searchedNodesCount should be 2 after second increment");
    }

    @ParameterizedTest
    @CsvSource({
        "0, 0, 0, 0, true",
        "10, 20, 10, 20, true",
        "10, 20, 10, 10, false",
        "10, 20, 20, 20, false",
        "10, 20, 20, 10, false"
    })
    void testEquals(long leaf1, long searched1, long leaf2, long searched2, boolean expected) {
        // Create first PerftResult
        PerftResult result1 = createPerftResult(leaf1, searched1);
        
        // Create second PerftResult
        PerftResult result2 = createPerftResult(leaf2, searched2);
        
        // Test equals
        assertEquals(expected, result1.equals(result2), "equals() returned unexpected result");
        
        // Test equals with null
        assertFalse(result1.equals(null), "equals(null) should return false");
        
        // Test equals with different class
        assertFalse(result1.equals("not a PerftResult"), "equals() with different class should return false");
        
        // Test reflexive
        assertEquals(result1, result1, "equals() should be reflexive");
        
        // Test symmetric
        if (expected) {
            assertEquals(result2, result1, "equals() should be symmetric");
        } else {
            assertNotEquals(result2, result1, "equals() should be symmetric");
        }
    }

    @ParameterizedTest
    @CsvSource({
        "0, 0, 0, 0, true",
        "10, 20, 10, 20, true",
        "10, 20, 10, 10, false",
        "10, 20, 20, 20, false"
    })
    void testHashCode(long leaf1, long searched1, long leaf2, long searched2, boolean expectedSameHash) {
        // Create first PerftResult
        PerftResult result1 = createPerftResult(leaf1, searched1);
        
        // Create second PerftResult
        PerftResult result2 = createPerftResult(leaf2, searched2);
        
        if (expectedSameHash) {
            assertEquals(result1.hashCode(), result2.hashCode(), 
                "hashCode() should be equal for equal objects");
        } else {
            // Note: Not strictly required for hash codes to be different for unequal objects,
            // but it's good practice to test for common cases
            assertNotEquals(result1.hashCode(), result2.hashCode(),
                "hashCode() should be different for different objects");
        }
    }

    @Test
    void testEqualsAndHashCodeConsistency() {
        PerftResult result1 = createPerftResult(10, 20);
        PerftResult result2 = createPerftResult(10, 20);
        
        // If two objects are equal, their hash codes must be equal
        if (result1.equals(result2)) {
            assertEquals(result1.hashCode(), result2.hashCode(),
                "Equal objects must have equal hash codes");
        }
    }

    // Helper method to create a PerftResult with specific values
    private PerftResult createPerftResult(long leafNodes, long searchedNodes) {
        PerftResult result = new PerftResult();
        result.setLeafNodesCount(leafNodes);
        for (long i = 0; i < searchedNodes; i++) {
            result.incrementSearchedNodesCount();
        }
        return result;
    }
}
