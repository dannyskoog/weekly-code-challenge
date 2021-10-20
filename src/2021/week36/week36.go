/**
 * Code challenge week 36, 2021 (language: GO, playground: https://play.golang.org/)
 *
 * Description:
 *     The code below is supposed to print a message of what the current second is. It is using a third party library to get the current
 *     second, which means that it’s not in control over said function. But there seems to be a problem! Sometimes when running the code
 *     nothing is printed. Say what?! One would at least expect the default message to be printed as a last resort!
 *
 * Questions:
 *     1. Why is it not always printing a message?
 *     2. How can it be adjusted to print a message at all times? (remember that we’re not in control over the third party library function)
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Current second is: N/A"

	msg = fmt.Sprintf("Current second is: %d", getCurrentSecond())

	fmt.Println(msg)
}

// Third party library function
func getCurrentSecond() int {
	currentSecond := time.Now().Second()
	isEvenSecond := currentSecond%2 == 0

	if isEvenSecond {
		panic(fmt.Errorf("Even seconds are evil!"))
	}

	return currentSecond
}
