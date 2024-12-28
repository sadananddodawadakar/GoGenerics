package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Summable interface {
	constraints.Ordered | string
}

func Add[T Summable](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(3, 7))         // 10
	fmt.Println(Add("Go", "lang")) // Golang
}
