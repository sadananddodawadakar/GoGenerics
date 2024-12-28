package main

import "fmt"

type Pair[T any, U any] struct {
	First  T
	Second U
}

func main() {
	p := Pair[int, string]{First: 1, Second: "one"}
	fmt.Println(p) // {1 one}

	q := Pair[int, int]{First: 1, Second: 10}
	fmt.Println(q)
}
