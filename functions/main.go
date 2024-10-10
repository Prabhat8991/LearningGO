package main

import "fmt"

func main() {
	fmt.Println(adder(3, 2))
	_, val := proAdder(3, 2, 3, 4)
	fmt.Println(val)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}
	return total, "Another value"
}
