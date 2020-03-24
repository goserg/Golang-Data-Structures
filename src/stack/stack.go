package stack

import (
	"fmt"
)

// Stack is LIFO (last in, first out) data structure
type Stack struct {
	data []interface{}
	size int
}

// NewStack is a constructor for the Stack
func NewStack() *Stack {
	s := new(Stack)
	s.data = []interface{}{}
	s.size = 0
	return s
}

// Push appends x to Data
// Time complexity O(1)
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
	s.size++
}

// Peek returns x from top of the stack
// Time complexity O(1)
func (s *Stack) Peek() (interface{}, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// Pop removes and returns x from top of the stack
// Time complexity O(1)
func (s *Stack) Pop() (interface{}, error) {
	if s.size == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.size--
	return x, nil
}

// Search for a data in the stack and get its distance from the top.
// This method starts the count of the position from 1.
// Return -1 if data not in the stack.
// Time complexity O(n)
func (s *Stack) Search(x interface{}) int {
	for i := 1; i <= len(s.data); i++ {
		if s.data[len(s.data)-i] == x {
			return i
		}
	}
	return -1
}

// Clear removes all data from the stack.
// Time complexity O(1)
func (s *Stack) Clear() {
	s.data = []interface{}{}
	s.size = 0
}

// IsEmpty returns false if the stack is empty else true.
// Time complexity O(1)
func (s *Stack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func (s Stack) String() string {
	return fmt.Sprint(s.data)
}

// GetSize returns size of the stack
// Time complexity O(1)
func (s *Stack) GetSize() int {
	return s.size
}
