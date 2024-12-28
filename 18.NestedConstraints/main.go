package main

import "fmt"

type Pair[T any] struct {
	First  T
	Second T
}

type NestedPair[T any, U any] struct {
	Inner Pair[T]
	Value U
}

func main() {
	nested := NestedPair[int, string]{
		Inner: Pair[int]{First: 10, Second: 20},
		Value: "Nested Value",
	}
	fmt.Println(nested)
}
