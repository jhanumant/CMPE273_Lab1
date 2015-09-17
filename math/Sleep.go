package math
import "time"

func modifiedSleep(secDurations int){
	<- time.After(time.Second * time.Duration(secDurations))
}

