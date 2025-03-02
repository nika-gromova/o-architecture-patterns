package collections

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(data T) {
	s.items = append(s.items, data)
}

func (s *Stack[T]) Pop() T {
	var res T
	if s.IsEmpty() {
		return res
	}
	res = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return res
}

func (s *Stack[T]) Top() (T, error) {
	var res T
	if s.IsEmpty() {
		return res, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) ToSlice() []T {
	return s.items
}
