package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"` //Removes this field from json
	Tags     []string
}

func main() {
	fmt.Println("Welcome to encoding Json")
	//EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"React js Bootcamp", 299, "lco.in", "abc123", []string{"webdev", "js"}},
		{"Go lang", 299, "lco.in", "abc123", []string{"golang", "js"}},
		{"Python", 299, "lco.in", "abc123", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
                "coursename": "React js Bootcamp",
                "Price": 299,
                "website": "lco.in",
                "Tags": ["webdev", "js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid..")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was NOT valid..")
	}

	// some cases where you want to add data to key value

	var myOnlineData map[string]interface{} //interface because we dont know about value
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

}
