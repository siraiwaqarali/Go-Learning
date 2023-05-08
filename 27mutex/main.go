package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	fmt.Println("Welcome to Goroutines")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, website := range websiteList {
		wg.Add(1) // Add 1 to the waitgroup counter
		go getStatusCode(website)
	}

	wg.Wait() // Wait for all the goroutines to finish
	fmt.Println("Signals:", signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() // Decrement the waitgroup counter

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in enpoint")
	} else {
		defer response.Body.Close()

		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()

		fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
	}

}
