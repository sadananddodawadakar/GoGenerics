package main

import "fmt"

func Reduce[T any](data []T, reducer func(T, T) T, initial T) T {
	result := initial
	for _, v := range data {
		result = reducer(result, v)
	}
	return result
}

func main() {
	sum := Reduce([]int{1, 2, 3}, func(a, b int) int { return a + b }, 0)
	fmt.Println(sum)
}
