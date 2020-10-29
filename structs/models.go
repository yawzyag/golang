package structs

// Rectangle struct definition
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Shape definition interface
type Shape interface {
	Area() float64
}
