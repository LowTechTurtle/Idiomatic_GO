package main

import (
	"fmt"
	"math"
)

// T implement interface stringer
type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)
	if d1 < d2 {
		return pair1
	} else {
		return pair2
	}
}

type Point2D struct {
	x, y int
}

func (p2d Point2D) String() string {
	return fmt.Sprintf("x: %d, y: %d", p2d.x, p2d.y)
}

func (p2d Point2D) Diff(from Point2D) float64 {
	X := p2d.x - from.x
	Y := p2d.y - from.y
	return math.Sqrt(float64(X*X) + float64(Y*Y))
}

type Point3D struct {
	X, Y, Z int
}

func (p3 Point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.X, p3.Y, p3.Z)
}

func (p3 Point3D) Diff(from Point3D) float64 {
	x := p3.X - from.X
	y := p3.Y - from.Y
	z := p3.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

func main() {
	d1 := FindCloser(Pair[Point2D]{Point2D{5, 15}, Point2D{10, 20}},
		Pair[Point2D]{Point2D{12, 13}, Point2D{10, 14}})
	fmt.Println(d1)

	d2 := FindCloser(Pair[Point3D]{Point3D{1, 2, 5}, Point3D{0, 1, 9}},
		Pair[Point3D]{Point3D{3, 2, 1}, Point3D{1, 2, 3}})
	fmt.Println(d2)
}
