package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func FilterStruct[T any](data []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range data {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Eve", 20},
	}
	adults := FilterStruct(people, func(p Person) bool { return p.Age >= 25 })
	fmt.Println(adults) // [{Alice 25} {Bob 30}]
}
