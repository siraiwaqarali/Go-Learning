package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "Welcome to time study"
	fmt.Println(welcome)

	presentTime := time.Now()
	fmt.Println("Present time is:", presentTime)
	fmt.Println("Formatted present time:", presentTime.Format("01-02-2006"))
	fmt.Println("Formatted present time:", presentTime.Format("01-02-2006 Monday"))
	fmt.Println("Formatted present time:", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2000, time.February, 23, 03, 30, 0, 0, time.UTC)
	fmt.Println("Created date is:", createdDate)
	fmt.Println("Formatted create date:", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
