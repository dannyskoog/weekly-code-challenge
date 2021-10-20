package main

import "fmt"

var _ Dog = (*goldenRetriever)(nil)

type goldenRetriever struct {
}

func (goldenRetriever) Bark() {
	fmt.Println("Bark")
}

func main() {
	goldenRetriever := goldenRetriever{}

	MakeAnimalSound(goldenRetriever)
}

// Third party library code below
type Dog interface {
	Bark()
}

type Cow interface {
	Moo()
}

func MakeAnimalSound(animal interface{}) {
	switch v := animal.(type) {
	case Dog:
		v.Bark()
	case Cow:
		v.Moo()
	default:
		fmt.Println("Don't know which sound to make")
	}
}
