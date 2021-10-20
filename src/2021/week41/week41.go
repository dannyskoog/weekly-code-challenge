/**
 * Code challenge week 41, 2021 (language: GO, playground: https://play.golang.org/)
 *
 * Description:
 *     The code below is supposed to print the sound of a dog (a Golden Retriever specifically). Everything is compiling perfectly
 *     and we therefore deploy it. But once hitting production we experience some unexpected printing (“Don’t know which sound to make”).
 *     And after some quick troubleshooting we realise that we’ve misspelled the function Barkk of the goldenRetriever struct (it should be
 *     just Bark). So we rapidly deploy a patch which resolves the issue (now printing “Bark”). But we’re still a bit shaken by how easily
 *     the typo was able to reach production. What if it happens again?!
 *
 * Questions:
 *     1. How could the code (that we own) be adjusted so that it doesn’t happen again?
 */

package main

import "fmt"

type goldenRetriever struct {
}

func (goldenRetriever) Barkk() {
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
