package main

import "fmt"

func main() {
	welcome := "Welcome to methods in golang"
	fmt.Println(welcome)

	waqar := User{"Waqar", "waqar@gmail.com", true, 23}
	fmt.Println(waqar)
	fmt.Printf("waqar details are: %+v\n", waqar)
	fmt.Println("Name is", waqar.Name)
	fmt.Println("Email is", waqar.Email)
	fmt.Println("Status is", waqar.Status)
	fmt.Println("Age is", waqar.Age)
	waqar.GetStatus()
	waqar.NewMail()
	fmt.Println(waqar)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "waqarupdated@gmail.com"
	fmt.Println("Email of user is changed to", u.Email)
}
