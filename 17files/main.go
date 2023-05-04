package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "This needs to go in a file - CodeWithWaqar"
	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile(file.Name())
}

func readFile(filePath string) {
	dataBytes, err := os.ReadFile(filePath)
	checkNilError(err)

	fmt.Println("File data is:")
	fmt.Println(string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err) // exit the program and show the error
	}
}
