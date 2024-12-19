package main

import (
	"cmp"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{f: f}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}

	if f(v, n.val) <= -1 {
		n.left = n.left.Add(f, v)
	} else {
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}

	if f(v, n.val) == 0 {
		return true
	}

	if f(v, n.val) <= -1 {
		return n.left.Contains(f, v)
	} else {
		return n.right.Contains(f, v)
	}
	return true
}

type Person struct {
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

func main() {
	t1 := NewTree(cmp.Compare[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

	t2 := NewTree(OrderPeople)
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Bob", 50})
	fmt.Println(t2.Contains(Person{"Bob", 30}))
	fmt.Println(t2.Contains(Person{"Fred", 25}))
}
