package main

import "fmt"

func Find[T comparable](data []T, target T) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(Find([]int{1, 2, 3}, 2))       // true
	fmt.Println(Find([]string{"a", "b"}, "z")) // false
}
