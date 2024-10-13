// **Basic Interface Definition** Interface Implementation**  Interface Variables**
package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (d Rectangle) Area() float64 {
	return d.Width * d.Height
}

func (d Rectangle) Perimeter() float64 {
	return 2 * (d.Width + d.Height)
}

func main() {
	var s Shape
	s = Rectangle{Width: 5, Height: 10}
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
