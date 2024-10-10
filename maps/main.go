package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS is short for ", languages["JS"])

	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
