package object

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		width:  10.0,
		height: 10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea1(t *testing.T) {
	type args struct {
		shape Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"area of shape", args{shape: Rectangle{10.0, 10.0}}, 100.0},
		{"area of circle", args{shape: Circle{10.0}}, 314.1592653589793},
		{"area of a triangle", args{shape: Triangle{12, 6}}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.shape.Area(); got != tt.want {
				t.Errorf("using %#v Area() = %g, want %g", tt.args.shape, got, tt.want)
			}
		})
	}
}
