package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course_name"`
	Price    int    `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`
	// Tags     []string `json:"tags"`
	Tags []string `json:"tags,omitempty"` // if tags is nil then it will not be shown in json
}

func main() {
	fmt.Println("Welcome to Jsons in Golang")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	var coursesList = []course{
		{"Go", 299, "CodeWithWaqar", "abc123", []string{"Backend Dev", "Golang"}},
		{"Python", 199, "CodeWithWaqar", "def123", nil},
		{"Flutter", 399, "CodeWithWaqar", "ghi123", []string{"Mobile Dev", "Dart"}},
	}

	// jsonData, err := json.Marshal(coursesList)
	jsonData, err := json.MarshalIndent(coursesList, "", "\t")
	handleErrNil(err)

	fmt.Printf("%s\n", jsonData)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
        "course_name": "Go",
        "price": 299,
        "website": "CodeWithWaqar",
        "tags": [
            "Backend Dev",
            "Golang"
        ]
    }
	`)

	var courseObj course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is  valid")
		err := json.Unmarshal(jsonDataFromWeb, &courseObj)
		handleErrNil(err)
		fmt.Printf("Course Obj: %#v\n", courseObj)

		// some cases where we just want to add data to key value pair
		var myOnlineData map[string]interface{}
		err = json.Unmarshal(jsonDataFromWeb, &myOnlineData)
		handleErrNil(err)
		fmt.Printf("My Online Data: %#v\n", myOnlineData)

		for k, v := range myOnlineData {
			fmt.Printf("For key %v, value is %v of type %T\n", k, v, v)
		}
	} else {
		fmt.Println("Json is not valid")
	}

}

func handleErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
