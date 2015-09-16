package main
import "fmt"

func fibonacciSeries(num int) int{
	if(num<0){
		panic("Please enter a positive integer")
	}
	if(num==0){
		return 0
	}
	if(num==1){
		return 1
	}
	return fibonacciSeries(num-1) + fibonacciSeries(num-2)
}

func main(){
	var num int
	fmt.Println("Enter n for F(n):")
	fmt.Scanf("%d\n",&num)
	series:= fibonacciSeries(num)
	fmt.Println("The answer:",series)
}
