package main

import "fmt"

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

func main() {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	fmt.Println(s.Contains(1)) // true
	fmt.Println(s.Contains(3)) // false
}
