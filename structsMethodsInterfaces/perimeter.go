package perimeter

//Rectangle new type for Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

//Perimeter returns perimeter of a shape given 2 sides
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

//Area returns area of a shape given 2 sides
func Area(rectangle Rectangle) float64 {
	return (rectangle.Height * rectangle.Width)
}
