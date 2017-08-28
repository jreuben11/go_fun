package main

import "fmt"
import "math"

// a basic interface
type geometry interface {
	area() float64
	perim() float64
}

// structs to implement this interface against
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// implement all methods in the interface for each struct
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// generic function with interface param
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The struct types both implement the interface - pass instances of these structs as arguments
	measure(r)
	measure(c)
}
