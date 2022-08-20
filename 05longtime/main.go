package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time on Go")
	timeNow := time.Now()
	fmt.Println(timeNow)
	fmt.Println(timeNow.Format("01-02-2006 15:04:05 Monday"))

	myDate := time.Date(1996, time.October, 22, 12, 0, 0, 0, time.Local)
	fmt.Println(myDate)
	fmt.Println(myDate.Format("02-01-2006 Monday"))
}
