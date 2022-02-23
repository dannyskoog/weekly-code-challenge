/**
* Code challenge week 8, 2022 (language/tool: GO, playground: https://play.golang.org/)
*
* Description:
*     The code below is supposed to print the description of a car. But once running the code we're surprised to see 'I am a Vehicle' get printed.
*     This is very disappointing since we were so happy being able to utilise the embedded 'Vehicle' struct's 'GetDescription' method.
*
* Questions:
*	  1. Why does 'I am a Vehicle' get printed?
*     2. How could the code be adjusted so that 'I am a Car' is printed?
 */

package main

import "fmt"

type Vehicle struct {
}

func (b *Vehicle) Type() string {
	return "Vehicle"
}

func (b *Vehicle) GetDescription() string {
	return fmt.Sprintf("I am a %s", b.Type())
}

type Car struct {
	Vehicle
}

func (b *Car) Type() string {
	return "Car"
}

func main() {
	car := &Car{}

	fmt.Println(car.GetDescription())
}
