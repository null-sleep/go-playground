package main

import (
	"fmt"
)

// Square blah blah
type Square struct {
	side int
}

// Rectangle blah blah
type Rectangle struct {
	height int
	width  int
}

// Shape blah blah
type Shape interface {
	Area() int
}

// Area blah blah
func (r Rectangle) Area() int {
	return r.height * r.width
}

// Area blah blah
func (s Square) Area() int {
	return s.side * s.side
}

// TripleArea blah blah
func TripleArea(s Shape) int {
	return s.Area() * 3
}

// Cant do this
// func TripleArea(r Rectangle) int {
// 	return r.Area() * 3
// }

func main() {
	fmt.Println("Hello, World!")
	square := Square{5}
	rectangle := Rectangle{2, 3}
	fmt.Printf("Area of square: %d\n", square.Area())
	fmt.Printf("Area of rectangle: %d\n", rectangle.Area())

	fmt.Printf("Area of square: %d\n", TripleArea(square))
	fmt.Printf("Area of rectangle: %d\n", TripleArea(rectangle))
}
