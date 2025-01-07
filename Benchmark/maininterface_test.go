package main

import (
	"testing"
)

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func BenchmarkInterfaceQueue(b *testing.B) {
	queue := Queue{}
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
		queue.Dequeue()
	}
}
