package main

import (
	"fmt"
	"math"
)
// here is a basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

//for our example well implement this interface on rect and circle types
type rect struct {
	width,height float64
}
type circle struct {
	radius float64
}

// to implement an interface in Go,
//we just need to implement all the methods in the interface
//here we do it for rectangles
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
//here we do it for circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2*math.Pi*c.radius
}

//if a variable has an interface type,
// then we can call methods that are in the named interface
//here's a generic measure function taking advantage of this 
//to work on any geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3,height: 4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}