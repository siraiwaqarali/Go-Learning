package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl = "http://localhost:3000"

func main() {
	fmt.Println("Welcome to web verb video")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	myurl := fmt.Sprintf("%v/postform", baseUrl)

	var data = url.Values{}
	data.Add("name", "Waqar Ali Siyal")
	data.Add("age", "23")

	response, err := http.PostForm(myurl, data)
	handleErrNil(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)

	content, err := io.ReadAll(response.Body)
	handleErrNil(err)

	fmt.Println("Content:", string(content))
}

func PerformPostJsonRequest() {
	myurl := fmt.Sprintf("%v/post", baseUrl)

	requestBody := strings.NewReader(`{
		"name": "Waqar Ali Siyal",
		"age": 23
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)
	handleErrNil(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	content, err := io.ReadAll(response.Body)
	handleErrNil(err)

	fmt.Println("Content:", string(content))
}

func PerformGetRequest() {
	myurl := fmt.Sprintf("%v/get", baseUrl)
	response, err := http.Get(myurl)
	handleErrNil(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, err := io.ReadAll(response.Body)
	handleErrNil(err)

	// fmt.Println("Content:", string(content))
	var responseString strings.Builder
	byteCount, err := responseString.Write(content)
	handleErrNil(err)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Content:", responseString.String())
}

func handleErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
