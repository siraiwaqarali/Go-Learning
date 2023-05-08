package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Goroutines")

	go greeter("Hello")
	greeter("World")
}

func greeter(s string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}
