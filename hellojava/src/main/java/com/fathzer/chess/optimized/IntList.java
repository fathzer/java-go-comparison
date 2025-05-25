package com.fathzer.chess.optimized;

/**
 * A resizable array implementation for primitive int values.
 * This class provides similar functionality to ArrayList<Integer> but avoids autoboxing overhead.
 */
public class IntList {
    private static final int DEFAULT_CAPACITY = 64;
    private int[] elements;
    private int size;

    /**
     * Constructs an empty list with an initial capacity of ten.
     */
    public IntList() {
        this.elements = new int[DEFAULT_CAPACITY];
        this.size = 0;
    }
    
    /**
     * Constructs a new IntList containing the elements of the specified IntList.
     * The new list will be a shallow copy of the original list.
     *
     * @param other the IntList whose elements are to be placed into this list
     * @throws NullPointerException if the specified IntList is null
     */
    public IntList(IntList other) {
        this.size = other.size;
        this.elements = new int[Math.max(size, DEFAULT_CAPACITY)];
        System.arraycopy(other.elements, 0, this.elements, 0, size);
    }

    /**
     * Appends the specified element to the end of this list.
     *
     * @param element element to be appended to this list
     */
    public void add(int element) {
        ensureCapacity(size + 1);
        elements[size++] = element;
    }
    
    /**
     * Appends all of the elements in the specified collection to the end of
     * this list, in the order that they are returned by the specified
     * collection's iterator.
     *
     * @param other collection containing elements to be added to this list
     * @throws NullPointerException if the specified collection is null
     */
    public void addAll(IntList other) {
        int numNew = other.size;
        ensureCapacity(size + numNew);
        System.arraycopy(other.elements, 0, elements, size, numNew);
        size += numNew;
    }

    /**
     * Removes and returns the last element from this list.
     *
     * @return the element that was removed from the list
     * @throws IllegalStateException if the list is empty
     */
    public int removeLast() {
        return elements[--size];
    }

    /**
     * Returns the element at the specified position in this list.
     *
     * @param index index of the element to return
     * @return the element at the specified position in this list
     * @throws IndexOutOfBoundsException if the index is out of range (index < 0 || index >= size())
     */
    public int get(int index) {
        return elements[index];
    }

    /**
     * Returns the number of elements in this list.
     *
     * @return the number of elements in this list
     */
    public int size() {
        return size;
    }

    /**
     * Removes all of the elements from this list. The list will be empty after this call returns.
     */
    public void clear() {
        size = 0;
    }
    
    /**
     * Returns true if this list contains the specified element.
     *
     * @param value element whose presence in this list is to be tested
     * @return true if this list contains the specified element, false otherwise
     */
    public boolean contains(int value) {
        for (int i = 0; i < size; i++) {
            if (elements[i] == value) {
                return true;
            }
        }
        return false;
    }

    /**
     * Increases the capacity of this IntList instance, if necessary, to ensure
     * that it can hold at least the number of elements specified by the minimum
     * capacity argument.
     *
     * @param minCapacity the desired minimum capacity
     */
    private void ensureCapacity(int minCapacity) {
        if (minCapacity > elements.length) {
            int newCapacity = elements.length + (elements.length >> 1); // Grow by 50%
            if (newCapacity < minCapacity) {
                newCapacity = minCapacity;
            }
            int[] newElements = new int[newCapacity];
            System.arraycopy(elements, 0, newElements, 0, size);
            elements = newElements;
        }
    }
}
