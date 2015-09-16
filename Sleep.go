package main
import "fmt"
import "time"

func modifiedSleep(secDurations int){
	<- time.After(time.Second * time.Duration(secDurations))
}

func main(){

	var secDurations int
	fmt.Println("Enter the sleep duration(in seconds)")
	fmt.Scanf("%d\n",&secDurations)
	modifiedSleep(secDurations)
	fmt.Println("Awake after sleeping for",secDurations)
}
