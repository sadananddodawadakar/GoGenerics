# GoGenerics

This repository contains examples of using generics in Go. The specific example provided here is an implementation of a generic linked list.

## Generic Linked List

This example demonstrates how to implement a generic linked list in Go. The `LinkedList` can store elements of any type.

### Implementation

The implementation consists of two main types:

- `Node[T]`: Represents a node in the linked list, where `T` is the type of the value stored in the node.
- `LinkedList[T]`: Represents the linked list itself.

### Usage

The `LinkedList` type provides two main methods:

- `Add(value T)`: Adds a new element to the end of the list.
- `Print()`: Prints all elements in the list.

### Example

Here is an example of how to use the `LinkedList`:

```go
package main

import "fmt"

type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T any] struct {
    Head *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
    newNode := &Node[T]{Value: value}
    if l.Head == nil {
        l.Head = newNode
    } else {
        current := l.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
}

func (l *LinkedList[T]) Print() {
    current := l.Head
    for current != nil {
        fmt.Println(current.Value)
        current = current.Next
    }
}

func main() {
    list := LinkedList[int]{}
    list.Add(10)
    list.Add(20)
    list.Add(30)
    list.Print()
}
