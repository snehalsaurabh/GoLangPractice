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
}