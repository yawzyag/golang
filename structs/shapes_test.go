package structs

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

func ExamplePerimeter() {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	fmt.Println(got)
	// Output: 40
}

func ExampleShape_Area() {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}
	fmt.Println(areaTests)
	// Output: [{Rectangle {12 6} 72} {Circle {10} 314.1592653589793} {Triangle {12 6} 36}]
}

func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}

func BenchmarkAreaRectangle(b *testing.B) {
	rectangle := Rectangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}

func BenchmarkAreaTriangle(b *testing.B) {
	triangle := Triangle{10.0, 10.0}
	for i := 0; i < b.N; i++ {
		triangle.Area()
	}
}

func BenchmarkAreaCircle(b *testing.B) {
	circle := Circle{10}
	for i := 0; i < b.N; i++ {
		circle.Area()
	}
}
