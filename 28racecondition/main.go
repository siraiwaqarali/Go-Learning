package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Race Condition")

	wg := &sync.WaitGroup{}
	// mutex := &sync.Mutex{}
	mutex := &sync.RWMutex{}

	var scores = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("One R")
		mutex.Lock()
		scores = append(scores, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Two R")
		mutex.Lock()
		scores = append(scores, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Three R")
		mutex.Lock()
		scores = append(scores, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)
	go func(wg *sync.WaitGroup, mutex *sync.RWMutex) {
		fmt.Println("Four R")
		mutex.RLock()
		fmt.Println("Scores (At Four R):", scores)
		mutex.RUnlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()

	fmt.Println("Scores (At End):", scores)
}
