package main

import "fmt"

type FloatyInt interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func double[T FloatyInt](x T) T {
	return 2 * x
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type Cats int

func (c Cats) String() string {
	return fmt.Sprintf("%d cats", c)
}

func Printy[T Printable](p T) {
	fmt.Println(p.String())
}

type Node[T comparable] struct {
	val  T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	root *Node[T]
}

func NewList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Add(v T) {
	if ll.root == nil {
		ll.root = &Node[T]{val: v}
	} else {
		ll.root.Add(v)
	}
}

func (n *Node[T]) Add(v T) {
	for n.next != nil {
		n = n.next
	}
	n.next = &Node[T]{val: v}
}

func (ll *LinkedList[T]) Index(v T) int {
	return ll.root.Index(v)
}

func (n *Node[T]) Index(v T) int {
	var i int
	for ; n != nil; n = n.next {
		if n.val == v {
			return i
		} else {
			i++
		}
	}
	return -1
}

func main() {
	var x uint32 = 122
	fmt.Println(double(x))
	var y float64 = 43215234.34252346
	fmt.Println(double(y))

	var cathouse Cats = 20
	Printy(cathouse)

	l := &LinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.root; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}

	l.Add(400)

	for curNode := l.root; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.val)
	}
}
