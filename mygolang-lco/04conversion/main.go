package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion in Golang")
	fmt.Println("Please rate the dish between 1 and 5")

	reader := bufio.NewReader(os.Stdin) 

	input, _ := reader.ReadString('\n') // The input is a string
	fmt.Println("Thank you for rating the dish as", input)

	// numRating = input + 1 - This will throw an error because you can't add a string to an integer

	// Type Conversion - Example 1
	numRating1, err := strconv.ParseFloat(input, 64) // This will convert the string to a float64 
	// Handling the error
	if err != nil {
		fmt.Println("An error occurred while converting the input to a number", err)
	} else {
		fmt.Println("Thank you for rating the dish as", numRating1 + 1)
	}
	// An error occurs while converting the input to a number. This is because when input is entered, it includes a newline character - \n ie 4\n for example, 

	
	// Type Conversion - Example 2
	numRating2, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // This will convert the string to a float64
	// Handling the error - This handles the error by removing the newline character
	if err != nil {
		fmt.Println("An error occurred while converting the input to a number", err)
	} else {
		fmt.Println("Thank you for rating the dish as", numRating2 + 1)
	}
}