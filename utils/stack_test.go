package utils

import (
	"testing"
)

// TestStackLength calls Stack.Length
// checking if the length of the stack is correct
func TestStackLength(t *testing.T) {
	test := Stack[int]{}

	size := 10

	for i := 0; i < size; i++ {
		test.Push(i)
	}

	if test.Length() != size {
		t.Errorf("The length is not correct : Wanted %v Got %v", size, test.Length())
	}
}

// TestStackPopEmpty calls Stack.Pop
// checking if the pop panic if the stack is empty
func TestStackPopEmpty(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The pop did not panic")
		}
	}()

	test := Stack[int]{}
	test.Pop()
}

// TestStackPopValue calls Stack.Pop
// checking if the value popped is correct
func TestStackPopValue(t *testing.T) {
	test := Stack[int]{}

	size := 10

	for i := 0; i < size; i++ {
		test.Push(i)
	}

	for i := size - 1; i >= 0; i-- {
		n := test.Pop()

		if n != i {
			t.Errorf("The pop value is not correct : Wanted %v Got %v", n, i)
		}
	}

}
