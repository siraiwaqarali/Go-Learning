package main

import "fmt"

func main() {
	// No inheritance in golang; No super or parent
	welcome := "Welcome to structs in golang"
	fmt.Println(welcome)

	waqar := User{"Waqar", "waqar@gmail.com", true, 23}
	fmt.Println(waqar)
	fmt.Printf("waqar details are: %+v\n", waqar)
	fmt.Println("Name is", waqar.Name)
	fmt.Println("Email is", waqar.Email)
	fmt.Println("Status is", waqar.Status)
	fmt.Println("Age is", waqar.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
