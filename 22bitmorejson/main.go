package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "xyz123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package data as json
	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Println("Json data:", string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var lcoCourses course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("Json is not valid")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is: %v and Value is: %v and Type is: %T\n", k, v, v)
	}
}
