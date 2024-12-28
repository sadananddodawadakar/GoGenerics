package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](data []T, target T) int {
	low, high := 0, len(data)-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] < target {
			low = mid + 1
		} else if data[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(BinarySearch(nums, 5)) // 2
}
