package services

import "fmt"

type HanderSecond struct{}

func (handerSecond *HanderSecond) Start() {
	fmt.Println("HanderSecond Start")
}

func (handerSecond *HanderSecond) Pause() {
	fmt.Println("HanderSecond Pause")
}

func (handerSecond *HanderSecond) Resume() {
	fmt.Println("HanderSecond Resume")
}

func (handerSecond *HanderSecond) Stop() {
	fmt.Println("HanderSecond Stop")
}
