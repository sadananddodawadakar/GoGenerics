package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value
}

func main() {
	stack := Stack[string]{}
	stack.Push("Go")
	stack.Push("Generics")
	fmt.Println(stack.Pop()) // Generics
	fmt.Println(stack.Pop()) // Go
}
