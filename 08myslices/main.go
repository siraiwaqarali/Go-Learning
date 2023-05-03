package main

import (
	"fmt"
	"sort"
)

func main() {
	welcome := "Welcome to a class on slices"
	fmt.Println(welcome)

	var fruitList []string = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of Fruit List: %T\n", fruitList)
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Length of Fruit List:", len(fruitList))

	fruitList = append(fruitList, "Banana", "Pineapple")
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Length of Fruit List:", len(fruitList))

	// Starts from index 1 and goes to the end
	// fruitList = append(fruitList[1:])
	// fmt.Println("Fruit List:", fruitList)
	// fmt.Println("Length of Fruit List:", len(fruitList))

	// Starts from index 1 and goes to index 3 (not inclusive)
	// fruitList = append(fruitList[1:3])
	// fmt.Println("Fruit List:", fruitList)
	// fmt.Println("Length of Fruit List:", len(fruitList))

	// Starts from index 0 and goes to index 3 (not inclusive)
	// fruitList = append(fruitList[:3])
	// fmt.Println("Fruit List:", fruitList)
	// fmt.Println("Length of Fruit List:", len(fruitList))

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777    // This will throw an error
	highScores = append(highScores, 555, 666, 777) // This will work and reallocate the memory
	fmt.Println("High Scores:", highScores)
	fmt.Println("Ints are sorted?", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("High Scores after sorting:", highScores)
	fmt.Println("Ints are sorted?", sort.IntsAreSorted(highScores))

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses:", courses)
	fmt.Println("Length of Courses:", len(courses))

	// Remove value from slice by index
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses:", courses)
	fmt.Println("Length of Courses:", len(courses))

}
