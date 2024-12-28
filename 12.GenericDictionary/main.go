package main

import "fmt"

type Dictionary[K comparable, V any] struct {
	items map[K]V
}

func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{items: make(map[K]V)}
}

func main() {
	dict := NewDictionary[string, int]()
	dict.items["age"] = 30
	fmt.Println(dict.items["age"]) // 30

	dict1 := NewDictionary[string, string]()
	dict1.items["Karnataka"] = "Bangalore"
	fmt.Println(dict1.items["Karnataka"])
}
