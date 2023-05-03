package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days of the week are: ", days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("For index %v, Value is %v\n", index, day)
	}

	rougeValue := 1
	for rougeValue < 10 {
		if rougeValue == 3 {
			rougeValue++
			continue
		}
		if rougeValue == 5 {
			goto codewithwaqar
		}
		if rougeValue == 8 {
			break
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

codewithwaqar:
	fmt.Println("Jumped to codewithwaqar")
}
