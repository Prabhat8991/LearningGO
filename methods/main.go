package main

import "fmt"

func main() {
	prabhat := User{"Prabhat Singh", "singh.prabhat@gmail.com", true, 32}
	fmt.Println(prabhat.GetStatus())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() bool {
	return u.Status
}
