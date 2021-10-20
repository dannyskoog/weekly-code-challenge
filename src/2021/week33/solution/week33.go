package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	str := "Initial value"

	go func() {
		time.Sleep(2 * time.Second)
		str = "Altered value"
		ch <- true
	}()

	<-ch
	time.Sleep(1 * time.Second)
	fmt.Println(str)
}
