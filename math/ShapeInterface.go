package math
import "math"


type Shape interface{
	area() float64
	perimeter() float64
}

type circle struct{
	radius float64
}

type rectangle struct {
	length float64
	width float64
}

func (rect rectangle) area() float64{
	return rect.length * rect.width
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (rect rectangle) perimeter() float64{
	return 2 * (rect.length + rect.width)
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func calculateArea(shape Shape) float64{
	return shape.area()
}

func calculatePerimeter(shape Shape) float64{
	return shape.perimeter()
}
