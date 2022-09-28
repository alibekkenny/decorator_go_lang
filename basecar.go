package main

import "fmt"

type BaseCar struct {
	Car        Car
	Price      int
	HorsePower int
	ModelName  string
}

func (baseCar *BaseCar) GetModelName() string {
	return baseCar.ModelName
}

func (baseCar *BaseCar) StartEngine() {
	fmt.Println(baseCar.GetModelName(), " engine starts with a weak sound!")
}

func (baseCar *BaseCar) Move() {
	fmt.Println(baseCar.ModelName, " moves slow!")
}

func (baseCar *BaseCar) GetHorsePower() int {
	return baseCar.HorsePower
}

func (baseCar *BaseCar) GetPrice() int {
	return baseCar.Price
}

func (baseCar *BaseCar) Construct() {
	fmt.Println(baseCar.GetModelName(), " has constructed as basic car. Horse power = ", baseCar.GetHorsePower(), "; price = ", baseCar.GetPrice())
}
