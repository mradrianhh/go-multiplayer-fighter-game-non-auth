package navigation

import (
	"errors"
)

// Stack ..
type Stack []Screen

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack.
func (s *Stack) Push(screen Screen) {
	*s = append(*s, screen)
}

// Pop removes and returns the top element of the stack.
func (s *Stack) Pop() (Screen, error) {
	if s.IsEmpty() {
		return Empty{}, errors.New("stack is empty")
	}

	index := len(*s) - 1
	screen := (*s)[index]
	*s = (*s)[:index]
	return screen, nil
}

// Peek returns the top element of the stack without removing it.
func (s *Stack) Peek() (Screen, error) {
	if s.IsEmpty() {
		return Empty{}, errors.New("stack is empty")
	}

	index := len(*s) - 1
	screen := (*s)[index]
	return screen, nil
}

// Len returns the length of the stack.
func (s *Stack) Len() int {
	return len(*s)
}
