package main

import "fmt"

type Comparable[T any] interface {
	Compare(other T) int
}

type Number int

func (n Number) Compare(other Number) int {
	return int(n - other)
}

func Max[T Comparable[T]](a, b T) T {
	if a.Compare(b) > 0 {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(Number(10), Number(20))) // 20
}
