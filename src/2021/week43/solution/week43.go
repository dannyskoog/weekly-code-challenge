package main

import "fmt"

type CarType string

const (
	Gasoline CarType = "Gasoline"
	Electric         = "Electric"
)

type Car struct {
	ID                int
	Brand             string
	Type              CarType
	BatteryPercentage *int
}

func NewCar(id int, brand string, carType CarType) *Car {
	car := &Car{
		ID:    id,
		Brand: brand,
		Type:  carType,
	}

	if carType == Electric {
		car.BatteryPercentage = new(int)
	}

	return car
}

func main() {
	car1 := NewCar(1, "BMW", Electric)
	car2 := NewCar(2, "Toyota", Gasoline)
	car3 := NewCar(3, "Hyundai", Electric)
	cars := []*Car{car1, car2, car3}

	for _, car := range cars {
		chargeBattery(car)
	}

	drainBattery(cars[2])

	for _, car := range cars {
		status := getBatteryPercentageStatus(car)
		fmt.Printf("Car: %d, Battery percentage status: %s\n", car.ID, status)
	}
}

func chargeBattery(car *Car) *Car {
	if car.Type == Electric {
		*car.BatteryPercentage = 100
	}

	return car
}

func drainBattery(car *Car) *Car {
	if car.Type == Electric {
		*car.BatteryPercentage = 0
	}

	return car
}

func getBatteryPercentageStatus(car *Car) string {
	if car.BatteryPercentage == nil {
		return "No battery to charge"
	}

	return fmt.Sprintf("%d%% charged", *car.BatteryPercentage)
}
