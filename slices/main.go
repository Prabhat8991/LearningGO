package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Banana", "Peach"}
	fruitList = append(fruitList, "papaya")
	fmt.Println("Printing Fruitlist ", fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println("Printing sliced Fruitlist ", fruitList)

	fruitList = append(fruitList[1:3]) // starts O based, last range is not inclusive also 0 based
	fmt.Println("Printing sliced Fruitlist ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 956
	highScores[2] = 123
	highScores[3] = 324
	//highScores[4] = 777 // Out of bound

	highScores = append(highScores, 555, 666, 321) // not out of bound

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "JS", "Java", "GO", "RESTFUL"}
	fmt.Println(courses)
	var index int = 2 // remove element at 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
