package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	spawnedGoroutineCount := 0
	var mutex = &sync.Mutex{}

	// Loop that spawns 1-10 goroutines
	for i := 1; i <= rand.Intn(10-1)+1; i++ {
		go func(id int) {
			mutex.Lock()
			newCount := spawnedGoroutineCount + 1
			time.Sleep(10 * time.Millisecond)
			spawnedGoroutineCount = newCount
			mutex.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Printf("Spawned goroutines count: %d\n", spawnedGoroutineCount)
}
