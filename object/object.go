package object

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.height + rectangle.width) * 2
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}

type Rectangle struct {
	height float64
	width  float64
}
