package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Sat"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for _, day := range days {
	// 	fmt.Println(day)
	// }

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			goto batman
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

batman:
	fmt.Println("Jumped to batman.....")
}
