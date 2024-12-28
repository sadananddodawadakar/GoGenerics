package main

import "fmt"

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q.Dequeue(), q.Dequeue()) // 1

	p := Queue[string]{}
	p.Enqueue("Go")
	p.Enqueue("Generics")
	fmt.Println(p.Dequeue(), p.Dequeue())
}
