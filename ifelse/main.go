package main

import "fmt"

func main() {
	loginCount := 3

	if loginCount > 4 {
		fmt.Println("Login passed")
	} else {
		fmt.Println("Login failed")
	}
}
