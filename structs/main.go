package main

import "fmt"

func main() {
	//There is no inheritence in Golang
	prabhat := User{"Prabhat", "singh.prabhat@gmail.com", true, 32}
	fmt.Println(prabhat)
	fmt.Printf("Prabhat's details are: %+v\n", prabhat)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
