package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/siraiwaqarali/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to MongoDB with GoLang!")
	fmt.Println("Starting Application on port 3000...")

	log.Fatal(http.ListenAndServe(":3000", router.Router()))
}
