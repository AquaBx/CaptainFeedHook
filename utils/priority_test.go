package utils

import (
	"container/heap"
	"testing"
)

// TestPriorityLength calls Priority.Len
// checking if the length of the Priority list is correct
func TestPriorityLength(t *testing.T) {
	test := Priority{}
	heap.Init(&test)

	size := 10

	for i := 0; i < size; i++ {
		heap.Push(&test, PriorityItem{Id: "item", NextCall: int64(i), LastCall: 0})
	}

	if test.Len() != 10 {
		t.Errorf("The length is not correct : Wanted %v Got %v", size, test.Len())
	}
}

// TestPriorityPop calls Priority.Pop
// checking if Priority list is correctly sorted
func TestPriorityPop(t *testing.T) {
	test := Priority{}
	heap.Init(&test)

	size := 10

	for i := 0; i < size; i++ {
		heap.Push(&test, PriorityItem{Id: "item", NextCall: int64(size - i - 1), LastCall: 0})
	}

	for i := 0; i < size; i++ {
		v := heap.Pop(&test).(PriorityItem)
		if v.NextCall != int64(i) {
			t.Errorf("The value is not correct : Wanted %v Got %v", i, v.NextCall)
		}
	}
}
