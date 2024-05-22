package main

import "fmt"

func main() {
	fmt.Println("If-else in Golang")

	loginCount := 23

	var result string // Declaring a variable result of type string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out! You are a VIP"
	} else {
		result = "Exactly 10 logins"
	}

	fmt.Println("Result:", result)

	// Short statement - Declare a variable inside if condition
	if num := 3; num < 10 {
		fmt.Println(num, "is less than 10")
	} else {	
		fmt.Println(num, "is not less than 10")
	}
	
	if err := 0; err != nil {
		fmt.Println(err)
	}
}