package main

import "fmt"

type Graph[T comparable] struct {
	adjacencyList map[T][]T
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{adjacencyList: make(map[T][]T)}
}

func (g *Graph[T]) AddEdge(from, to T) {
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
}

func main() {
	graph := NewGraph[string]()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("A", "D")
	graph.AddEdge("B", "A")
	graph.AddEdge("C", "A")
	graph.AddEdge("D", "A")
	fmt.Println(graph.adjacencyList)
}
