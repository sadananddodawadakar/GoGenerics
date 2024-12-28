package main

import "fmt"

func Reverse[T any](data []T) []T {
	n := len(data)
	for i := 0; i < n/2; i++ {
		data[i], data[n-i-1] = data[n-i-1], data[i]
	}
	return data
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(Reverse(nums)) // [4 3 2 1]
}
