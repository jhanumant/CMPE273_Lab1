package math

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
