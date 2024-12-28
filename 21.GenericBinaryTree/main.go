package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Ordered] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (n *TreeNode[T]) Insert(value T) {
	if value == n.Value {
		return
	} else if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode[T]{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode[T]{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *TreeNode[T]) Print() {
	if n.Left != nil {
		n.Left.Print()
	}
	fmt.Println(n.Value)
	if n.Right == nil {
		return
	}
	n.Right.Print()

}

func main() {
	root := &TreeNode[int]{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(8)
	root.Insert(3)
	root.Insert(12)
	root.Print()
}
