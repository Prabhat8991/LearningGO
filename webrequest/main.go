package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com/"

func main() {
	response, error := http.Get(url)

	if error != nil {
		panic(error)
	}

	fmt.Println("Response.... ", response)

	defer response.Body.Close()

	databytes, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	fmt.Println("Data bytes values ", string(databytes))

}
