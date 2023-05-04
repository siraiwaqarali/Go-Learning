package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests in Golang")

	response, err := http.Get(url)
	handleErrNil(err)
	defer response.Body.Close() // Caller's responsibility to close the connection

	fmt.Printf("Response is of type: %T\n", response)

	dataBytes, err := io.ReadAll(response.Body)
	handleErrNil(err)

	content := string(dataBytes)
	fmt.Println(content)
}

func handleErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
