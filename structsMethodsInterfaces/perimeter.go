package perimeter

import "math"

//Rectangle new type for rectangle info
type Rectangle struct {
	Width  float64
	Height float64
}

//Perimeter Perim of a rectangle given two sides
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

//Area Area of a rectangle given two sides
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle new type for circle info
type Circle struct {
	Radius float64
}

//Area Area of a circle given struct's radius
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Triangle Area of a triangle given a side
type Triangle struct {
	Side float64
}

//Area Area of a circle given struct's radius
func (t Triangle) Area() float64 {
	return (t.Side * t.Side) / 2
}
