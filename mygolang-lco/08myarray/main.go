package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang")
     
	// Declaring an array
	var myArray [5]int
	fmt.Println("Array:", myArray) // This will print [0 0 0 0 0] as the array is not initialized

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50

	fmt.Println("Array:", myArray) // This will print [10 20 30 40 50] as the array is initialized

	// Declaring and initializing an array
	var myArray1 = [5]int{10, 20, 30, 40, 50} 
	fmt.Println("Array1:", myArray1) // This will print [10 20 30 40 50] as the array is initialized

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[3] = "Mango"
	fmt.Println("Fruits:", fruits) // This will print [Apple Banana  Mango] as the array is initialized with empty strings
	fmt.Println("Length of Fruits:", len(fruits)) // This will print 4 as the length of the array is 4 even though the 2nd element is empty

	var vegetables = [3]string{"Carrot", "Beans", "Spinach"} 
	fmt.Println("Vegetables:", vegetables) // This will print [Carrot Beans Spinach] as the array is initialized
}
