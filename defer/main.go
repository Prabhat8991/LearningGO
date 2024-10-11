package main

import "fmt"

func main() {
	//defer takes the statement to last and executes in LIFO manner
	defer fmt.Println("World1")
	defer fmt.Println("World2")
	defer fmt.Println("World3")
	fmt.Println("Hello")
}
