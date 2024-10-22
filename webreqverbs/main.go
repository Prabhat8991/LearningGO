package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//PerformGetRequest()
	//PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println("Response content ", string(content))

}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
     {
		"Name":"Prabhat",
		"Designation":"Lead Con"
	 }
   `)
	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	resp, err := ioutil.ReadAll(response.Body)
	fmt.Println("Response... ", string(resp))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}

	data.Add("firstName", "Prabhat")
	data.Add("Designation", "Lead Con")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	resp, _ := ioutil.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Println("Response...", string(resp))

}
