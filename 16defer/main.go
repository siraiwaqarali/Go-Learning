package main

import "fmt"

// defer is used to delay the execution of a statement until function exits.
// defer is often used where e.g. ensure and finally would be used in other languages.
// If a function has multiple defer statements, they are added on to a stack
// and executed in Last In First Out (LIFO) order.

func main() {
	fmt.Println("Welcome to defer in Golang")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
