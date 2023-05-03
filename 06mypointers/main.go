package main

import "fmt"

func main() {
	welcome := "Welcome to a class on pointers"
	fmt.Println(welcome)

	var ptr *int
	fmt.Println("Default value of pointer is:", ptr)

	myNumber := 23
	var myNumberPtr *int = &myNumber
	fmt.Println("Value of myNumberPtr is:", myNumberPtr)
	fmt.Println("Value in myNumberPtr is:", *myNumberPtr)

	*myNumberPtr = *myNumberPtr * 2
	fmt.Println("New value (from myNumberPtr):", *myNumberPtr)
	fmt.Println("New value (from myNumber):", myNumber)
}
