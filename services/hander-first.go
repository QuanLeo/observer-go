package services

import "fmt"

type HanderFirst struct{}

func (handerFirst *HanderFirst) Start() {
	fmt.Println("HanderFirst Start")
}

func (handerFirst *HanderFirst) Pause() {
	fmt.Println("HanderFirst Pause")
}

func (handerFirst *HanderFirst) Resume() {
	fmt.Println("HanderFirst Resume")
}

func (handerFirst *HanderFirst) Stop() {
	fmt.Println("HanderFirst Stop")
}
