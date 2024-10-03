package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text to print")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
}
