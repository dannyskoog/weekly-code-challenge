package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Car struct {
	Brand string
	Color string
}

func main() {
	var cars = make([]Car, 10)

	car := Car{
		Brand: "Tesla",
		Color: "Red",
	}

	addCarAtRandomIndex(cars, car)

	foundIndex := -1

	for i, car := range cars {
		if (car != Car{}) {
			foundIndex = i
			break
		}
	}

	fmt.Printf("Found index: %d\n", foundIndex)
}

func addCarAtRandomIndex(cars []Car, car Car) {
	length := len(cars)

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(length) // Note: Always generates the same number in the GO Playground

	cars[index] = car
}
