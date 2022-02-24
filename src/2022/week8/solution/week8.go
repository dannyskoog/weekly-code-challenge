package main

import "fmt"

type TypeGetter interface {
	Type() string
}

type Vehicle struct {
}

func (b *Vehicle) Type() string {
	return "Vehicle"
}

func (b *Vehicle) GetDescription(typeGetter TypeGetter) string {
	typeStr := b.Type()

	if _, ok := typeGetter.(TypeGetter); ok {
		typeStr = typeGetter.Type()
	}

	return fmt.Sprintf("I am a %s", typeStr)
}

type Car struct {
	Vehicle
}

func (c *Car) Type() string {
	return "Car"
}

func (c *Car) GetDescription() string {
	return c.Vehicle.GetDescription(c)
}

func main() {
	car := &Car{}

	fmt.Println(car.GetDescription())
}
