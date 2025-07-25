package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 7}

	fmt.Println("Rectangle:")
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())

	fmt.Println("\nCircle:")
	fmt.Println("Area:", circle.Area())
	fmt.Println("Perimeter:", circle.Perimeter())
}
