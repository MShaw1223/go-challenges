package objects

import "math"

type Shape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	Width int
	Height int
}

type Circle struct {
	Radius float32
}

func (r * Rectangle) Area() int {
	return r.Width * r.Height
}

func (c * Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (r * Rectangle) Perimeter()int {
	return 2*(r.Width + r.Height)
}

func (c * Circle) Perimeter()float32{
	return 2 * math.Pi * c.Radius
}