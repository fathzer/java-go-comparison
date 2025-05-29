// Package optimized provides optimized data structures for chess engine implementation.
package optimized

// DefaultCapacity is the default initial capacity of an IntList.
const DefaultCapacity = 64

// IntList is a resizable array implementation for primitive int values.
// It provides similar functionality to a slice but with a more controlled API.
type IntList struct {
	elements []int
	size     int
}

// NewIntList creates a new empty IntList with the default capacity.
func NewIntList() *IntList {
	return &IntList{
		elements: make([]int, DefaultCapacity),
		size:     0,
	}
}

// NewIntListFrom creates a new IntList that is a copy of another IntList.
func NewIntListFrom(other *IntList) *IntList {
	if other == nil {
		panic("IntList: source cannot be nil")
	}

	// Create a new list with the same size as the other list
	newList := &IntList{
		elements: make([]int, max(other.size, DefaultCapacity)),
		size:     other.size,
	}
	// Copy the elements
	copy(newList.elements, other.elements[:other.size])
	return newList
}

// Add appends the specified element to the end of this list.
func (l *IntList) Add(element int) {
	l.ensureCapacity(l.size + 1)
	l.elements[l.size] = element
	l.size++
}

// AddAll appends all elements from another IntList to this list.
func (l *IntList) AddAll(other *IntList) {
	if other == nil {
		return
	}

	numNew := other.size
	l.ensureCapacity(l.size + numNew)
	copy(l.elements[l.size:], other.elements[:numNew])
	l.size += numNew
}

// RemoveLast removes and returns the last element from this list.
func (l *IntList) RemoveLast() int {
	if l.size == 0 {
		panic("IntList: cannot remove from empty list")
	}

	l.size--
	return l.elements[l.size]
}

// Get returns the element at the specified position in this list.
func (l *IntList) Get(index int) int {
	if index < 0 || index >= l.size {
		panic("IntList: index out of bounds")
	}
	return l.elements[index]
}

// Size returns the number of elements in this list.
func (l *IntList) Size() int {
	return l.size
}

// Clear removes all elements from this list.
func (l *IntList) Clear() {
	l.size = 0
}

// Contains returns true if this list contains the specified value.
func (l *IntList) Contains(value int) bool {
	for i := 0; i < l.size; i++ {
		if l.elements[i] == value {
			return true
		}
	}
	return false
}

// ensureCapacity ensures that the list can hold at least minCapacity elements.
func (l *IntList) ensureCapacity(minCapacity int) {
	if minCapacity > len(l.elements) {
		newCapacity := len(l.elements) + (len(l.elements) >> 1) // Grow by 50%
		if newCapacity < minCapacity {
			newCapacity = minCapacity
		}

		newElements := make([]int, newCapacity)
		copy(newElements, l.elements[:l.size])
		l.elements = newElements
	}
}

// ToSlice returns a new slice containing all elements in this list.
func (l *IntList) ToSlice() []int {
	slice := make([]int, l.size)
	copy(slice, l.elements[:l.size])
	return slice
}
