package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Current second is: N/A"

	defer defaultPrint(msg)

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

func defaultPrint(msg string) {
	if err := recover(); err != nil {
		fmt.Println(msg)
	}
}
