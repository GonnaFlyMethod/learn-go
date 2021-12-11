package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40},
		{name: "Circle", shape: Circle{Radius: 3}, want: 18.84955592153876},
		{name: "Triangle", shape: Triangle{Base: 5.0, Height: 3.0, A: 2.0, B: 3.0, C: 2.3}, want: 7.3},
	}

	for _, tt := range perimeterTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()

			if got != tt.want {
				t.Errorf("%#v | got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 5.0, Height: 3.0, A: 2.0, B: 3.0, C: 2.3}, hasArea: 7.5},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v | actualArea: %g, wantedArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
