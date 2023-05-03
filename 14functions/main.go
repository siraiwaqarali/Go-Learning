package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	fmt.Println("Result is:", adder(3, 5))
	fmt.Println("Result is:", proAdder(6, 2, 8))
	fmt.Println("Result is:", proAdder(1, 5, 4, 2))
	fmt.Println("Result is:", proAdder(4, 7, 4, 6, 9, 2, 6))

	name, age := multipleReturnValues()
	fmt.Printf("Name is %v and age is %v\n", name, age)
}

func greeter() {
	fmt.Println("Hello from Golang")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func multipleReturnValues() (string, int) {
	return "Waqar Ali Siyal", 23
}
