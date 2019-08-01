package main

import (
	"fmt"
)

func main() {
	// Array
	//var fruitArr [2]string

	// Assign Values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	fruitArr := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape"}
	// Print Array
	fmt.Println(fruitArr)

	// Print Array @ index
	fmt.Println(fruitArr[1])

	// Print Slice
	fmt.Println(fruitSlice)

	// Print Length of Array & Slice
	fmt.Println(len(fruitArr), len(fruitSlice))

	// Print Range of Slice
	fmt.Println(fruitSlice[1:3])
}
