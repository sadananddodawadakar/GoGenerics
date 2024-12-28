package main

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

func SortDescending[T constraints.Ordered](data []T) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
}

func main() {
	nums := []int{5, 3, 9, 1}
	SortDescending(nums)
	fmt.Println(nums) // [9 5 3 1]
	str := []string{"Cat", "Apple", "Ball", "Elephant", "Dog"}
	SortDescending(str)
	fmt.Println(str)
}
