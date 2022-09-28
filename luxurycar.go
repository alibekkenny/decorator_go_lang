package main

import "fmt"

type LuxuryCar struct {
	Car          Car
	ComfortLevel string
	Interior     string
}

func (luxuryCar *LuxuryCar) GetModelName() string {
	return luxuryCar.Car.GetModelName()
}

func (luxuryCar *LuxuryCar) Move() {
	fmt.Println(luxuryCar.GetModelName(), " moves comfortably!")
}

func (luxuryCar *LuxuryCar) StartEngine() {
	fmt.Println(luxuryCar.GetModelName(), " engine starts with nice sound!")
}

func (luxuryCar *LuxuryCar) GetHorsePower() int {
	actualHorsePower := luxuryCar.Car.GetHorsePower()
	return actualHorsePower + 100
}

func (luxuryCar *LuxuryCar) GetPrice() int {
	actualPrice := luxuryCar.Car.GetPrice()
	return actualPrice + 200000
}

func (luxuryCar *LuxuryCar) Construct() {
	fmt.Println(luxuryCar.GetModelName(), " has constructed as luxury car. Horse power = ", luxuryCar.GetHorsePower(), "; price = ", luxuryCar.GetPrice())
}
