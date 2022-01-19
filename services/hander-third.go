package services

import "fmt"

type HanderThird struct{}

func (handerThird *HanderThird) Start() {
	fmt.Println("HanderThird Start")
}

func (handerThird *HanderThird) Pause() {
	fmt.Println("HanderThird Pause")
}

func (handerThird *HanderThird) Resume() {
	fmt.Println("HanderThird Resume")
}

func (handerThird *HanderThird) Stop() {
	fmt.Println("HanderThird Stop")
}
