package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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
}

func getStatusCode(endpoint string) {
	defer wg.Done() // Decrement the waitgroup counter

	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in enpoint")
	}
	defer response.Body.Close()
	fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)

}
