package main

import "fmt"

type SportCar struct {
	Car           Car
	HandlingLevel string
	EngineType    string
}

func (sportCar *SportCar) GetModelName() string {
	return sportCar.Car.GetModelName()
}

func (sportCar *SportCar) Move() {
	fmt.Println(sportCar.GetModelName(), " moves fast!")
}

func (sportCar *SportCar) StartEngine() {
	fmt.Println(sportCar.GetModelName(), " engine starts with loud sound!")
}

func (sportCar *SportCar) GetHorsePower() int {
	actualHorsePower := sportCar.Car.GetHorsePower()
	return actualHorsePower + 300
}

func (sportCar *SportCar) GetPrice() int {
	actualPrice := sportCar.Car.GetPrice()
	return actualPrice + 100000
}

func (sportCar *SportCar) Construct() {
	fmt.Println(sportCar.GetModelName(), " has constructed as sport car. Horse power = ", sportCar.GetHorsePower(), "; price = ", sportCar.GetPrice())
}
