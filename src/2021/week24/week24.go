/**
 * Code challenge week 24, 2021 (language: GO, playground: https://play.golang.org/)
 *
 * Questions:
 *     1. What would you replace the question mark (?) with in order to print the index that the car was added at?
 */

package main

import (
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

	// ?
}

func addCarAtRandomIndex(cars []Car, car Car) {
	length := len(cars)

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(length) // Note: Always generates the same number in the GO Playground

	cars[index] = car
}
