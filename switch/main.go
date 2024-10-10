package main

import "fmt"

func main() {
	rank := 2

	switch rank {
	case 1:
		fmt.Println("Congrats on first rank")
	case 2:
		fmt.Println("Congrats on second rank")
		fallthrough
	case 3:
		fmt.Println("Congrats on third rank")
	default:
		fmt.Println("Enter correct rank")
	}
}
