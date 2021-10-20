/**
 * Code challenge week 33, 2021 (language: GO, playground: https://play.golang.org/)
 *
 * Description:
 *     The code below is supposed to print the value of the str variable.
 *
 * Questions:
 *     1. Why is it always printing “Initial value”?
 *     2. How can it be adjusted to print “Altered value”?
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	str := "Initial value"

	go func() {
		time.Sleep(2 * time.Second)
		str = "Altered value"
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(str)
}
