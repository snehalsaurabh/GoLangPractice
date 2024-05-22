package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in Golang")

	// Declaring a pointer
	var ptr *int
	fmt.Println("Pointer:", ptr) // This will print nil as the pointer is not pointing to any memory location

	myNumber := 20
	var ptr1 *int = &myNumber // & is used to get the memory location of the variable
	fmt.Println("Pointer1:", ptr1) // This will print the memory location of myNumber
	fmt.Println("Value at Pointer1:", *ptr1) // This will print the value at the memory location of myNumber

	*ptr1 = *ptr1*5 // Changing the value at the memory location of myNumber
	fmt.Println("New Value is:", myNumber)	
}
