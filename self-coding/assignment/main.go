package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func pr(sh shape) {
	fmt.Println(sh.getArea())
}

func main() {
	t := triangle{2, 4}
	s := square{13}

	pr(t)
	pr(s)
}
