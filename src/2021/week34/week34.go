/**
 * Code challenge week 34, 2021 (language: GO, playground: https://play.golang.org/)
 *
 * Description:
 *     The code below is supposed to print the count of spawned goroutines. But it seems to be always spawning just 1. Hum… That can’t be
 *     correct, right?!
 *
 * Questions:
 *     1. Why is it always printing that just 1 goroutine was spawned?
 *     2. How can it be adjusted to print the correct count of spawned goroutines?
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	spawnedGoroutineCount := 0

	// Loop that spawns 1-10 goroutines
	for i := 1; i <= rand.Intn(10-1)+1; i++ {
		go func(id int) {
			newCount := spawnedGoroutineCount + 1
			time.Sleep(10 * time.Millisecond)
			spawnedGoroutineCount = newCount
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Printf("Spawned goroutines count: %d\n", spawnedGoroutineCount)
}
