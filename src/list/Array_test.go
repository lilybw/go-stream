package list

import (
	"sync"
	"testing"
)

func TestNewArray(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)
	if arr.Size() != 0 {
		t.Errorf("NewArray should create an empty array, got size %d", arr.Size())
	}
}

func TestNewArrayWithCapacity(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr, _ := NewArrayWithCapacity[int](0, equals)
	if cap(arr.elements) != 0 {
		t.Errorf("Expected capacity 0, got %d", cap(arr.elements))
	}

	arrNegative, err := NewArrayWithCapacity[int](-1, equals)
	if err == nil {
		t.Errorf("An error must be returned for negative capacity")
	}
	if err != nil && arrNegative == nil {
		t.Errorf("Array must default to default implementation when an error occurs")
	}
}

func TestAddExtremeValues(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)
	arr.Push(0)
	arr.Push(int(^uint(0) >> 1)) // Maximum int value

	if !arr.Contains(0) || !arr.Contains(int(^uint(0)>>1)) {
		t.Errorf("Failed to add or verify extreme values")
	}
}

func TestAddRemove(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)
	arr.Push(1)
	arr.Push(2)
	if !arr.Contains(1) || !arr.Contains(2) {
		t.Errorf("Add operation failed to insert elements correctly")
	}
	if arr.Remove(2) == false || arr.Contains(2) {
		t.Errorf("Remove operation failed")
	}
	if arr.Size() != 1 {
		t.Errorf("Expected size 1, got %d", arr.Size())
	}
}

func TestSubListAndEquals(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	sub := arr.SubList(1, 3)
	if sub.Size() != 2 {
		t.Errorf("SubList size expected 2, got %d", sub.Size())
	}
	if !sub.Contains(2) || !sub.Contains(3) {
		t.Errorf("SubList missing elements")
	}
	cloned := arr.Clone()
	if !arr.Equals(cloned) {
		t.Errorf("Clone should be equal to the original")
	}
}

func TestFilter(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	filtered := arr.Filter(func(x int) bool { return x%2 == 0 })
	if filtered.Size() != 1 || !filtered.Contains(2) {
		t.Errorf("Filter failed, expected [2], got %s", filtered.ToString())
	}
}
func TestConcurrency(t *testing.T) {
	//TODO: Flakey, but then again, this is not a concurrent data structure
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			arr.Push(val)
			if !arr.Contains(val) {
				t.Errorf("Concurrency error: Added value %d not found", val)
			}
		}(i)
	}
	wg.Wait()

	if arr.Size() != 100 {
		t.Errorf("Concurrency error: Expected size 100, got %d", arr.Size())
	}
}

func TestOutOfRange(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr := NewArray(equals)
	arr.Push(1)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out-of-range access")
		}
	}()
	_ = arr.Get(5) // Should panic
}

func TestComplexEquality(t *testing.T) {
	equals := func(a, b int) bool { return a == b }
	arr1, _ := NewArrayWithCapacity[int](10, equals)
	arr2, _ := NewArrayWithCapacity[int](5, equals)

	arr1.Push(1)
	arr1.Push(2)
	arr2.Push(1)
	arr2.Push(2)

	if !arr1.Equals(arr2) {
		t.Errorf("Equals failed for arrays with same elements but different capacities")
	}

	arr1.Push(3)
	arr2.Push(3)
	arr1.Set(0, 3) // Change order
	arr1.Set(2, 1)

	if arr1.Equals(arr2) {
		t.Errorf("Equals should fail for arrays with same elements in different order")
	}
}
