package main

import (
	"fmt"
	"quanleo/observer"
	"quanleo/services"
	"time"
)

func main() {
	fmt.Println("Start")

	subjectMonitor := observer.SubjectMonitor{}

	handerFirst := services.HanderFirst{}
	handerSecond := services.HanderSecond{}
	handerThird := services.HanderThird{}

	subjectMonitor.Register(&handerFirst)
	subjectMonitor.Register(&handerSecond)
	subjectMonitor.Register(&handerThird)

	// for test
	subjectMonitor.Notification(observer.START)
	customeSleep()

	subjectMonitor.Notification(observer.STOP)
	customeSleep()

	subjectMonitor.Notification(observer.START)
	customeSleep()

	subjectMonitor.Notification(observer.RESUME)

	fmt.Println("End")
}

func customeSleep() {
	time.Sleep(2 * time.Second)
	fmt.Println("\n-------------")
}
