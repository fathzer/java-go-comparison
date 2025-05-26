package optimized

import (
	"testing"
)

func TestEmptyList(t *testing.T) {
	list := NewIntList()
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
}

func TestAddAndGet(t *testing.T) {
	list := NewIntList()
	list.Add(10)
	list.Add(20)
	list.Add(30)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}

	tests := []struct {
		index    int
		expected int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
	}

	for _, test := range tests {
		if got := list.Get(test.index); got != test.expected {
			t.Errorf("list.Get(%d) = %d, want %d", test.index, got, test.expected)
		}
	}
}

func TestRemoveLast(t *testing.T) {
	list := NewIntList()
	list.Add(10)
	list.Add(20)


	if last := list.RemoveLast(); last != 20 {
		t.Errorf("Expected last element 20, got %d", last)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}

	if last := list.RemoveLast(); last != 10 {
		t.Errorf("Expected last element 10, got %d", last)
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
}

func TestContains(t *testing.T) {
	list := NewIntList()
	if list.Contains(10) {
		t.Error("Empty list should not contain any elements")
	}

	list.Add(10)
	list.Add(20)
	list.Add(30)


	tests := []struct {
		value    int
		expected bool
	}{
		{10, true},
		{20, true},
		{30, true},
		{40, false},
	}

	for _, test := range tests {
		if got := list.Contains(test.value); got != test.expected {
			t.Errorf("list.Contains(%d) = %v, want %v", test.value, got, test.expected)
		}
	}

	// Test after removing elements
	list.RemoveLast() // Removes 30
	if list.Contains(30) {
		t.Error("List should not contain 30 after removal")
	}

	list.Clear()
	if list.Contains(10) || list.Contains(20) {
		t.Error("List should be empty after clear")
	}
}

func TestClear(t *testing.T) {
	list := NewIntList()
	list.Add(10)
	list.Add(20)

	list.Clear()

	if list.Size() != 0 {
		t.Errorf("Expected size 0 after clear, got %d", list.Size())
	}
}

func TestCopyConstructor(t *testing.T) {
	// Test with empty list
	emptyList := NewIntList()
	emptyCopy := NewIntListFrom(emptyList)
	if emptyCopy.Size() != 0 {
		t.Errorf("Expected size 0, got %d", emptyCopy.Size())
	}

	// Test with non-empty list
	list := NewIntList()
	list.Add(10)
	list.Add(20)
	list.Add(30)


	copiedList := NewIntListFrom(list)

	// Verify the copy has the same elements
	if list.Size() != copiedList.Size() {
		t.Errorf("Expected size %d, got %d", list.Size(), copiedList.Size())
	}

	for i := 0; i < list.Size(); i++ {
		if list.Get(i) != copiedList.Get(i) {
			t.Errorf("Element at index %d differs: %d vs %d", i, list.Get(i), copiedList.Get(i))
		}
	}

	// Verify modifying the original doesn't affect the copy
	copySize := copiedList.Size()
	list.Add(40)
	if copiedList.Size() != copySize {
		t.Error("Modifying original affected copy's size")
	}
	if copiedList.Contains(40) {
		t.Error("Modifying original affected copy's contents")
	}

	// Verify modifying the copy doesn't affect the original
	copiedList.RemoveLast()
	if list.Size() == copiedList.Size() {
		t.Error("Modifying copy affected original's size")
	}
	if !list.Contains(30) {
		t.Error("Modifying copy affected original's contents")
	}
}

func TestCopyConstructorNil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic with nil source")
		}
	}()

	_ = NewIntListFrom(nil)
}

func TestAddAll(t *testing.T) {
	// Test adding to empty list
	source := NewIntList()
	source.Add(1)
	source.Add(2)
	source.Add(3)

	target := NewIntList()
	target.AddAll(source)

	if target.Size() != 3 {
		t.Fatalf("Expected size 3, got %d", target.Size())
	}

	expected := []int{1, 2, 3}
	for i, v := range expected {
		if target.Get(i) != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, target.Get(i))
		}
	}

	// Test adding to non-empty list
	additional := NewIntList()
	additional.Add(4)
	additional.Add(5)

	target.AddAll(additional)


	if target.Size() != 5 {
		t.Fatalf("Expected size 5, got %d", target.Size())
	}

	expected = []int{1, 2, 3, 4, 5}
	for i, v := range expected {
		if target.Get(i) != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, target.Get(i))
		}
	}

	// Test adding empty list
	prevSize := target.Size()
	target.AddAll(NewIntList())
	if target.Size() != prevSize {
		t.Errorf("Expected size %d after adding empty list, got %d", prevSize, target.Size())
	}
}

func TestLargeLists(t *testing.T) {
	const size1 = 1000
	const size2 = 2000

	// Test AddAll with large lists
	source := NewIntList()
	for i := 0; i < size1; i++ {
		source.Add(i)
	}

	target := NewIntList()
	for i := 0; i < size2; i++ {
		target.Add(i * 2)
	}

	target.AddAll(source)

	// Verify the result
	if target.Size() != size1+size2 {
		t.Fatalf("Expected size %d, got %d", size1+size2, target.Size())
	}

	// Check first size2 elements (original target elements)
	for i := 0; i < size2; i++ {
		expected := i * 2
		if got := target.Get(i); got != expected {
			t.Fatalf("At index %d: expected %d, got %d", i, expected, got)
		}
	}

	// Check remaining elements (from source)
	for i := 0; i < size1; i++ {
		expected := i
		if got := target.Get(size2 + i); got != expected {
			t.Fatalf("At index %d: expected %d, got %d", size2+i, expected, got)
		}
	}

	// Test Contains with large list
	// The target list contains:
	// 1. Numbers from 0 to (size2-1)*2 (even numbers)
	// 2. Then numbers from 0 to size1-1
	// So all numbers from 0 to size1-1 should be in the list (from the source)
	// And all even numbers up to (size2-1)*2 should be in the list (from the target)
	for i := 0; i < size1; i++ {
		if !target.Contains(i) {
			t.Errorf("Expected to find %d in the list", i)
		}
	}

	// Check some even numbers from the target
	for i := 0; i < 100; i += 2 {
		if !target.Contains(i) {
			t.Errorf("Expected to find %d in the list", i)
		}
	}

	// Also check some numbers that shouldn't be in the list
	if target.Contains(-1) {
		t.Error("Did not expect to find -1 in the list")
	}
	// size2*2 would be 4000, but since we only added numbers up to (size2-1)*2 (3998) and size1-1 (999)
	// 3999 is not in the list (it's odd and > size1-1)
	if target.Contains(3999) {
		t.Error("Did not expect to find 3999 in the list")
	}
	// Check for a number that's definitely not in the list
	// The list contains numbers up to max(1998, 999) = 1998
	// So 1999 is not in the list (it's odd and > 999)
	if target.Contains(1999) {
		t.Error("Did not expect to find 1999 in the list")
	}
}
