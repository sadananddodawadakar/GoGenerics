package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Merge[T constraints.Ordered](a, b []T) []T {
	var result []T
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	fmt.Println(Merge(a, b)) // [1 2 3 4 5 6]
}
