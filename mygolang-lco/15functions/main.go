package main

import "fmt"

func main() { // This is the main function and acts as the entry point of the program
	fmt.Println("Welcome to Functions in Golang")
	greeter() // This is how you call a function in Go

	fmt.Println("- This is not allowed in Go. Can't define a function inside a function -")
	// func greeterTwo() { // This is a function named greeterTwo 
	// 	fmt.Println("Hi! I am a function too")
	// } // This function is not accessible outside the main function

	result := adder(3,5) // This is how you call a function with parameters in Go
	fmt.Println("Result is: ", result) 

	proRes, return2 := proAdder(1,2,3,4,5,6,7,8,9,10) // This is how you call a function with variable number of arguments in Go
	fmt.Println("Pro Result is: ", proRes)
	fmt.Println("Second return value: ", return2)
}

func greeter() { // This is a function named greeter
	fmt.Println("Hi! I am a function")
}

func adder (a int, b int) int { // This is a function named adder which takes two parameters a and b of type int and returns an int
	return a + b
}
// The int after the parenthesis (a int, b int) is the return type of the function - int in this case. It is called function signature


// This is a function named proAdder which takes a variable number of arguments of type int and returns an int
func proAdder(values ...int) (int, string) { // (int, string) is the return type of the function - int and string in this case - Multiple return values
	result := 0
	for _, value := range values {
		result += value
	}
	return result, "Example of returning two values together of two different types"
}