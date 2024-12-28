package main

import "fmt"

type Box[T any] struct {
	Value T
}

func (b Box[T]) GetValue() T {
	return b.Value
}

func main() {
	b := Box[int]{Value: 42}
	fmt.Println(b.GetValue()) // 42

	c := Box[string]{Value: "Hello"}
	fmt.Println(c.GetValue()) // Hello
}
