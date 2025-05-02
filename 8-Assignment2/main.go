package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	printArea() float64
}

func printArea(s shape) {
	fmt.Println(s.printArea())
}

func (t triangle) printArea() float64 {
	return (t.height * t.base) / 2
}

func (s square) printArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{height: 10, base: 5}
	s := square{sideLength: 4}

	printArea(t)
	printArea(s)
}