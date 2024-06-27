package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		area  float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, area: 100.0},
		{name: "Circle", shape: Circle{10.0}, area: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, area: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.area {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.area)
			}
		})
	}
}
