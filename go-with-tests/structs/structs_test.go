package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("got %#v %g want %.2f", shape, got, want)
		}
	}
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}
	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
