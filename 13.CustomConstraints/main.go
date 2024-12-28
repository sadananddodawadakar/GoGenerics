package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Sum(1, 2))     // 3
	fmt.Println(Sum(1.1, 2.2)) // 3.3
}
