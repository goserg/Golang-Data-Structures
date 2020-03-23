package stack

import (
	"fmt"
)

// Stack is LIFO (last in, first out) data structure
type Stack struct {
	Data []float32
	len  int
}

// Push appends x to Data
// Time complexity O(1)
func (s *Stack) Push(x float32) {
	s.Data = append(s.Data, x)
	s.len++
}

// Peek returns x from top of the stack
// Time complexity O(1)
func (s *Stack) Peek() (float32, error) {
	if s.len == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.Data[len(s.Data)-1], nil
}

// Pop removes and returns x from top of the stack
// Time complexity O(1)
func (s *Stack) Pop() (float32, error) {
	if s.len == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	x := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	s.len--
	return x, nil
}

// Search for a data in the stack and get its distance from the top.
// This method starts the count of the position from 1.
// Return -1 if data not in the stack.
// Time complexity O(n)
func (s *Stack) Search(x float32) int {
	for i := 1; i <= len(s.Data); i++ {
		if s.Data[len(s.Data)-i] == x {
			return i
		}
	}
	return -1
}

// Clear removes all data from the stack.
// Time complexity O(1)
func (s *Stack) Clear() {
	s.Data = []float32{}
	s.len = 0
}

// IsEmpty returns false if the stack is empty else true.
// Time complexity O(1)
func (s *Stack) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}

func (s Stack) String() string {
	return fmt.Sprint(s.Data)
}

// Size returns size of the stack
// Time complexity O(1)
func (s *Stack) Size() int {
	return s.len
}
