package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

func main() {
	fmt.Println("Welcome to handling urls in Golang")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)
	handleErrNil(err)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("The type of the queryParams is %T\n", queryParams)
	fmt.Println("Query Params:", queryParams)

	fmt.Println("Value for key coursename is", queryParams["coursename"])
	fmt.Println("Value for key paymentid is", queryParams["paymentid"])

	for _, val := range queryParams {
		fmt.Println("Param is", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=waqar",
	}

	urlFromPartsOfUrl := partsOfUrl.String()
	fmt.Println("Formed from parts of url:", urlFromPartsOfUrl)
}

func handleErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
