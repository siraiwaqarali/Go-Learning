package main

import "fmt"

func main() {
	welcome := "Welcome to maps in golang"
	fmt.Println(welcome)

	// var languages map[string]string = map[string]string{}
	// var languages map[string]string = map[string]string{"CS": "C#"}
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"
	languages["TS"] = "Typescript"
	languages["RB"] = "Ruby"

	fmt.Println("Languages:", languages)
	fmt.Println("Length of Languages:", len(languages))
	fmt.Println("JS shorts for:", languages["JS"])

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	delete(languages, "TS")
	fmt.Println("Languages:", languages)
	fmt.Println("Length of Languages:", len(languages))
}
