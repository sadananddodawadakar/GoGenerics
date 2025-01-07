package main

import (
	"testing"
)

type QueueG[T any] struct {
	items []T
}

func (q *QueueG[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *QueueG[T]) Dequeue() T {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func BenchmarkGenericQueue(b *testing.B) {
	queue := QueueG[int]{}
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
		queue.Dequeue()
	}
}
