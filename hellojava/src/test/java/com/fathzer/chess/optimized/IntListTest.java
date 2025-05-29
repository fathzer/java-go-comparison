package com.fathzer.chess.optimized;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.ValueSource;

class IntListTest {
    private IntList list;

    @BeforeEach
    void setUp() {
        list = new IntList();
    }

    @Test
    void testEmptyList() {
        assertEquals(0, list.size());
    }

    @Test
    void testAddAndGet() {
        list.add(10);
        list.add(20);
        list.add(30);
        
        assertEquals(3, list.size());
        assertEquals(10, list.get(0));
        assertEquals(20, list.get(1));
        assertEquals(30, list.get(2));
    }

    @Test
    void testRemoveLast() {
        list.add(10);
        list.add(20);
        
        assertEquals(20, list.removeLast());
        assertEquals(1, list.size());
        assertEquals(10, list.get(0));
        
        assertEquals(10, list.removeLast());
        assertEquals(0, list.size());
    }
    
    @Test
    void testContains() {
        assertFalse(list.contains(10));
        
        list.add(10);
        list.add(20);
        list.add(30);
        
        assertTrue(list.contains(10));
        assertTrue(list.contains(20));
        assertTrue(list.contains(30));
        assertFalse(list.contains(40));
        
        // Test after removing elements
        list.removeLast(); // Removes 30
        assertTrue(list.contains(10));
        assertTrue(list.contains(20));
        assertFalse(list.contains(30));
        
        list.clear();
        assertFalse(list.contains(10));
        assertFalse(list.contains(20));
    }
    
    @ParameterizedTest
    @ValueSource(ints = {1, 10, 100, 1000})
    void testContainsWithLargeList(int size) {
        // Fill the list with even numbers
        for (int i = 0; i < size; i++) {
            list.add(i * 2);
        }
        
        // Test that all even numbers are contained
        for (int i = 0; i < size; i++) {
            if (i % 2 == 0) {
                assertTrue(list.contains(i), "Should contain " + i);
            } else {
                assertFalse(list.contains(i), "Should not contain " + i);
            }
        }
    }

    @Test
    void testClear() {
        list.add(10);
        list.add(20);
        
        list.clear();
        
        assertEquals(0, list.size());
    }
    
    @Test
    void testCopyConstructor() {
        // Test with empty list
        IntList emptyCopy = new IntList(list);
        assertEquals(0, emptyCopy.size());
        
        // Test with non-empty list
        list.add(10);
        list.add(20);
        list.add(30);
        
        IntList copy = new IntList(list);
        
        // Verify the copy has the same elements
        assertEquals(list.size(), copy.size());
        for (int i = 0; i < list.size(); i++) {
            assertEquals(list.get(i), copy.get(i), "Element at index " + i + " differs");
        }
        
        // Verify modifying the original doesn't affect the copy
        list.add(40);
        assertNotEquals(list.size(), copy.size());
        assertFalse(copy.contains(40));
        
        // Verify modifying the copy doesn't affect the original
        copy.removeLast();
        assertNotEquals(list.size() - 1, copy.size());
        assertTrue(list.contains(30));
    }
    
    @Test
    void testCopyConstructorNull() {
        assertThrows(NullPointerException.class, () -> new IntList(null));
    }
    
    @Test
    void testCopyConstructorLargeList() {
        // Test with a list larger than DEFAULT_CAPACITY
        final int size = 1000;
        for (int i = 0; i < size; i++) {
            list.add(i);
        }
        
        IntList copy = new IntList(list);
        assertEquals(size, copy.size());
        for (int i = 0; i < size; i++) {
            assertEquals(i, copy.get(i), "Element at index " + i + " differs");
        }
    }

    @Test
    void testAddAll() {
        // Test adding to empty list
        IntList other = new IntList();
        other.add(1);
        other.add(2);
        other.add(3);
        
        list.addAll(other);
        
        assertEquals(3, list.size());
        assertEquals(1, list.get(0));
        assertEquals(2, list.get(1));
        assertEquals(3, list.get(2));
        
        // Test adding to non-empty list
        IntList another = new IntList();
        another.add(4);
        another.add(5);
        
        list.addAll(another);
        
        assertEquals(5, list.size());
        assertEquals(4, list.get(3));
        assertEquals(5, list.get(4));
        
        // Test adding empty list
        int prevSize = list.size();
        list.addAll(new IntList());
        assertEquals(prevSize, list.size());
    }
    
    @Test
    void testAddAllWithLargeLists() {
        // Test adding lists that require capacity expansion
        final int size1 = 1000;
        final int size2 = 2000;
        
        // Fill the source list
        IntList source = new IntList();
        for (int i = 0; i < size1; i++) {
            source.add(i);
        }
        
        // Fill the target list
        for (int i = 0; i < size2; i++) {
            list.add(i * 2);
        }
        
        // Add all elements from source to target
        list.addAll(source);
        
        // Verify the result
        assertEquals(size1 + size2, list.size());
        
        // Check original elements are still there
        for (int i = 0; i < size2; i++) {
            assertEquals(i * 2, list.get(i));
        }
        
        // Check added elements are correct
        for (int i = 0; i < size1; i++) {
            assertEquals(i, list.get(size2 + i));
        }
    }
    
    @Test
    void testAddAllNull() {
        assertThrows(NullPointerException.class, () -> list.addAll(null));
    }

    @Test
    void testGetOutOfBounds() {
        list.add(10);
        assertThrows(IndexOutOfBoundsException.class, () -> list.get(-1));
    }

    @Test
    void testAutoResize() {
        // Add more elements than initial capacity (10)
        for (int i = 0; i < 20; i++) {
            list.add(i);
        }
        
        assertEquals(20, list.size());
        for (int i = 0; i < 20; i++) {
            assertEquals(i, list.get(i));
        }
    }
}
