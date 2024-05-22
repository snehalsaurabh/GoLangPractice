package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	// Declaring a slice
	var fruits = []string{"Apple", "Banana", "Mango", "Pineapple", "Orange"} // This is a slice but you have to initialize it with values in this way
	fmt.Printf("Type of Fruitlist is %T\n:", fruits)                         // This will print []string as the type of the slice

	fruits = append(fruits, "Grapes") // This will append Grapes to the slice
	fmt.Println("Fruits:", fruits)    // This will print [Apple Banana Mango Pineapple Orange Grapes]

	fruits = append(fruits[1:])    // This will remove the element at 0th index from the slice
	fmt.Println("Fruits:", fruits) // This will print [Banana Mango Pineapple Orange Grapes]

	fruits = append(fruits[1:3])   // This will remove the elements at 0th and 3rd index onwards from the slice. The 3rd index is exclusive
	fmt.Println("Fruits:", fruits) // This will print [Mango Pineapple]

	// Declaring a slice with make function
	highscore := make([]int, 3) // This will create a slice of length 3 with default values
	fmt.Println("Highscore:", highscore) // This will print [0 0 0] as the slice is initialized with default values

	highscore[0] = 100
	highscore[1] = 200
	highscore[2] = 300
	fmt.Println("Highscore:", highscore) // This will print [100 200 300] as the slice is initialized

	// highscore[4] = 500 -> // This will throw an error as the index is out of range

	highscore = append(highscore, 500, 600) // This will append 500 and 600 to the slice
	// We expected it to throw an error as the length of the slice is 3 but it didn't. This is because the slice is dynamically resized.
	fmt.Println("Highscore:", highscore) // This will print [100 200 300 500 600] as the slice is resized and 500 and 600 are appended

	// Sorting a slice
	sort.Ints(highscore) // This will sort the slice in ascending order
	fmt.Println("Highscore:", highscore) // This will print [100 200 300 500 600] as the slice is sorted in ascending order
	fmt.Println(sort.IntsAreSorted(highscore)) // This will print true as the slice is sorted in ascending order or false if it was not sorted
}
