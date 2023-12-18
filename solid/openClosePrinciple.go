package solid

import "math"

// The Open/Closed Principle states that software entities (classes, modules, functions, etc.) should be open for extension but closed for modification. This principle encourages developers to write code that is flexible and can be extended without the need for significant modifications.

// Suppose we have a simple function that calculates the area of a rectangle:

type Rectangle struct {
	Width  float64
	Height float64
}

func Area1(rectangle *Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// If we need to add support for calculating the area of a circle, the current implementation would require modification:

type Circle struct {
	Radius float64
}

func Area2(shape interface{}) float64 {
	switch s := shape.(type) {
	case *Rectangle:
		return s.Width * s.Height
	case *Circle:
		return math.Pi * math.Pow(s.Radius, 2)
	default:
		return 0
	}
}

// To follow the Open/Closed Principle, we can define an interface and implement it for each shape:
type Shape interface {
	Area() float64
}

func (r *Rectangle) Area1() float64 {
	return r.Width * r.Height
}

func (c *Circle) Area1() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
