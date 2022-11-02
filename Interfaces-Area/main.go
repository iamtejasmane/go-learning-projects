package main

import (
	"fmt"
)

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

func main() {
	s := square{10.0}
	s.printArea()

	t := triangle{5.2, 4.4}
	t.printArea()

	printArea(s)
	printArea(t)

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) printArea() {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
