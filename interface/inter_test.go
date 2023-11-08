package inter

import (
	"testing"
)

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Perimeter(rectangle)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestAre(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	got := Area(rectangle)
// 	want := 72.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangels", func(t *testing.T) {
		r := Rectangle{12.0, 6.0}
		checkArea(t, r, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		checkArea(t, c, 314.1592653589793)
	})
}

func TestAreaInterface(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{width: 12, height: 6},
			hasArea: 72.0,
		},
		{
			name:    "Circle",
			shape:   Circle{Radius: 10},
			hasArea: 314.1592653589793,
		},
		{
			name:    "Triangle",
			shape:   Triangle{Side: 12, Height: 6},
			hasArea: 36.0,
		},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %.2f want %.2f", got, tt.hasArea)
			}
		})

	}
}
