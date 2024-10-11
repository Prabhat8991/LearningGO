package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Writing something into file...."

	file, error := os.Create("./myfile.txt")

	if error != nil {
		panic(error)
	}

	length, error := io.WriteString(file, content)

	if error != nil {
		panic(error)
	}

	fmt.Println("Length is:,", length)
	defer file.Close() // recommended way to use defer

	readFile("./myfile.txt")
}

func readFile(fileName string) {
	databyte, error := ioutil.ReadFile(fileName)

	if error != nil {
		panic(error)
	}

	fmt.Println("Text data inside file is", string(databyte))
}
