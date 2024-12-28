package main

import "fmt"

// Generic swap, that can swap any type in Go

func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	fmt.Println("Demo: Generic swap")
	x, y := 10, 20

	x, y = Swap(x, y)
	fmt.Println("Swapped values...", x, y)

	a, b := "Hi", "Hello"
	a, b = Swap(a, b)
	fmt.Println("Swapped values...", a, b)

	type Struct struct {
		a, b int
	}

	p, q := Struct{10, 20}, Struct{100, 200}

	p, q = Swap(p, q)
	fmt.Println("Swapped values...", p, q)

}
