package main

import (
	"fmt"
)

func main() {
	car := &BaseCar{ModelName: "Toyota Supra", HorsePower: 300, Price: 3000}
	car.Construct()
	car.StartEngine()
	car.Move()
	fmt.Println()

	offroadCar := &OffroadCar{
		Car: car,
	}
	offroadCar.Construct()
	offroadCar.StartEngine()
	offroadCar.Move()
	fmt.Println()

	sportCar := &SportCar{
		Car: offroadCar,
	}
	sportCar.Construct()
	sportCar.StartEngine()
	sportCar.Move()
	fmt.Println()

	luxuryCar := &LuxuryCar{
		Car: sportCar,
	}
	luxuryCar.Construct()
	luxuryCar.StartEngine()
	luxuryCar.Move()
	fmt.Println()
}
