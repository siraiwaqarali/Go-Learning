package main

import "fmt"

func main() {
	welcome := "Welcome to array in golang"
	fmt.Println(welcome)

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"
	fruitList[2] = "Grape"

	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Length of Fruit List:", len(fruitList))

	var vegList [3]string = [3]string{"Potato", "Tomato", "Brinjal"}
	fmt.Println("Veg List:", vegList)
	fmt.Println("Length of Veg List:", len(vegList))
}
