package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and Deadlock")

	// Create a channel
	channel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	// Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		value, isChannelOpen := <-ch
		// fmt.Println(<-ch)
		fmt.Println(isChannelOpen)
		fmt.Println(value)
		wg.Done()
	}(channel, wg)
	// Send only channel
	go func(ch chan<- int, wg *sync.WaitGroup) {

		ch <- 5
		ch <- 6
		close(ch)
		wg.Done()
	}(channel, wg)

	wg.Wait()
}
