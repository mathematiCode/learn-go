package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4, 5}
	got := Perimeter(rectangle)
	want := float64(18)

	if got != want {
		t.Errorf("got %.5f but wanted %.5f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 3, Height: 8}, expectedArea: 24.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 2, Height: 4}, expectedArea: 6},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g but wanted %g", got, want)
		}
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.expectedArea)
		})
	}
}

// t.Run("rectangles", func(t *testing.T) {
// 	rectangle := Rectangle{4, 5}
// 	checkArea(rectangle, 20)
// })

// t.Run("circles", func(t *testing.T) {
// 	circle := Circle{3}
// 	checkArea(circle, 9*math.Pi)
// })

// t.Run("triangles", func(t *testing.T) {
// 	triangle := Triangle{3, 4}
// 	checkArea(triangle, 6)
// })
