package main
import "fmt"
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

func calculateArea(shape Shape){
	fmt.Println("The area is:",shape.area())
}

func calculatePerimeter(shape Shape){
	fmt.Println("The perimeter is:",shape.perimeter())
}

func main(){

	c:= new(circle)
	rect:= new(rectangle)
	fmt.Println("Please enter your choice:\n 1) Area of Circle\n 2) Perimeter of Circle \n 3) Area of rectangle \n 4) Perimeter of rectangle")
	var choice int64 
	fmt.Scanf("%d\n",&choice)
	switch choice{
		case 1:
			fmt.Println("Enter the radius:")
			fmt.Scanf("%f\n",&c.radius)
			calculateArea(c)
			break
		case 2:
			fmt.Println("Enter the radius:")
			fmt.Scanf("%f\n",&c.radius)
			calculatePerimeter(c)
			break
		case 3:
			fmt.Println("Enter length of rectangle:")
			fmt.Scanf("%f\n",&rect.length)
			fmt.Println("Enter width of rectangle:")
			fmt.Scanf("%f\n",&rect.width)
			calculateArea(rect)
			break
		case 4:
			fmt.Println("Enter length of rectangle:")
			fmt.Scanf("%f\n",&rect.length)
			fmt.Println("Enter width of rectangle:")
			fmt.Scanf("%f\n",&rect.width)
			calculatePerimeter(rect)
			break
		default:
			fmt.Println("Please enter a correct choice.")
			break

	}
}
