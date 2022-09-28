package main

import "fmt"

type OffroadCar struct {
	Car       Car
	PassLevel string
	FrameType string
}

func (offroadCar *OffroadCar) GetModelName() string {
	return offroadCar.Car.GetModelName()
}

func (offroadCar *OffroadCar) Move() {
	fmt.Println(offroadCar.GetModelName(), " moves stubbornly!")
}

func (offroadCar *OffroadCar) StartEngine() {
	fmt.Println(offroadCar.GetModelName(), " engine starts with rough sound!")
}
func (offroadCar *OffroadCar) GetHorsePower() int {
	actualHorsePower := offroadCar.Car.GetHorsePower()
	return actualHorsePower + 200
}

func (offroadCar *OffroadCar) GetPrice() int {
	actualPrice := offroadCar.Car.GetPrice()
	return actualPrice + 50000
}

func (offroadCar *OffroadCar) Construct() {
	fmt.Println(offroadCar.GetModelName(), " has constructed as offroad car. Horse power = ", offroadCar.GetHorsePower(), "; price = ", offroadCar.GetPrice())
}
