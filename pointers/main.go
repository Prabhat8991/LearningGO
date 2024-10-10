package main

import "fmt"

func main() {

	var ptr *int //default value nil

	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr2 *int = &myNumber

	fmt.Println("Value of pointer2 is ", ptr2)  // Address that pointer is holding
	fmt.Println("Value of pointer2 is ", *ptr2) // Actual value

	*ptr = *ptr + 2
	fmt.Println("Value of myNumber is ", myNumber)

}
