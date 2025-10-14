package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (1.0 / 2.0) * t.Base * t.Height
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return rectangle.Width*2 + rectangle.Height*2
}
