package math
import "testing"
import "time"


type fibSeries struct{
	number int
	answer int
}

type rectArea struct{
	length float64
	width float64
	areaOfRect float64
	perimeterOfRect float64
}

type circleArea struct{
	radius float64
	areaOfCircle float64
	perimeterOfCircle float64
}

var outputCircle = [] circleArea{
	{2,12.566370614359172,12.566370614359172},
}

var outputRectangle =[] rectArea{
	{2,3,6,10},
}
var outputFibonacci= [] fibSeries{
	{1,1},
	{5,5},
	{4,3},
}

func TestFibonacci(t *testing.T){
	for _ , pair:= range outputFibonacci{
		ans:= fibonacciSeries(pair.number)
		if ans!= pair.answer{
			t.Error("For",pair.number,
					"The fibonacci is:",pair.answer,"but got:",ans)
		}
	}
}

func TestCircleArea(t *testing.T){
	c:= new(circle)
	for _,pair:= range outputCircle{
		c.radius=pair.radius
		ans := calculateArea(c)
		if(ans!=pair.areaOfCircle){
			t.Error("For",pair.radius,"The area is:",pair.areaOfCircle,"but got:",ans)
		}
	}
}

func TestCirclePerimeter(t *testing.T){
	c:= new(circle)
	for _,pair:= range outputCircle{
		c.radius=pair.radius
		ans := calculatePerimeter(c)
		if(ans!=pair.areaOfCircle){
			t.Error("For",pair.radius,"The perimeter is:",pair.perimeterOfCircle,"but got:",ans)
		}
	}
}

func TestRectangleArea(t *testing.T){
	rect := new(rectangle)
	for _,pair:= range outputRectangle{
		rect.length=pair.length
		rect.width=pair.width
		ans := calculateArea(rect)
		if(ans!=pair.areaOfRect){
			t.Error("For",pair.length,"&",pair.width,"The area is:",pair.areaOfRect,"but got:",ans)
		}
	}
}

func TestRectanglePerimeter(t *testing.T){
	rect := new(rectangle)
	for _,pair:= range outputRectangle{
		rect.length=pair.length
		rect.width=pair.width
		ans := calculatePerimeter(rect)
		if(ans!=pair.perimeterOfRect){
			t.Error("For",pair.length,"&",pair.width,"The perimeter is:",pair.perimeterOfRect,"but got:",ans)
		}
	}
}

func TestSleepTime(t *testing.T){
	sleepTime:=5
	waitingTime:= time.Now()
	modifiedSleep(sleepTime)
	waitedTime:= time.Now()
	if waitedTime.Sub(waitingTime) > time.Second * time.Duration(6){
		t.Error("The time before",sleepTime,"was",waitingTime, "but now it is",time.Now())
	}
}
