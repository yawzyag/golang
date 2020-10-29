package structs

import "math"

// Area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area of triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Perimeter of a shape
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
