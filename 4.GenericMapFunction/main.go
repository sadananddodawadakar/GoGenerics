package main

import "fmt"

func Map[T any, U any](data []T, mapper func(T) U) []U {
	var result []U
	for _, v := range data {
		result = append(result, mapper(v))
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	squares := Map(nums, func(n int) int { return n * n })
	fmt.Println(squares) // [1, 4, 9]
}
