package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // S and N denote public functions. Saving the sentence imported from the package bufio and os.
	fmt.Println("Enter the rating for our Pizza: ") // Print.ln injects a new line '\n' after the output.
	// fmt.Scanln(&rating) // Scanln is used to read the input from the user. It reads the input from the user and stores it in the memory location of the variable.

	// comma ok || err syntax
	input, _ := reader.ReadString('\n') // ReadString reads the input from the user and returns the input as a string. '\n' is the delimiter.
	// _ is used to ignore the error returned by the function. The error is ignored because we are not handling the error in this program.
	// we could have used err in the place of _ to handle the error. This structure is called comma ok syntax.

	// If you don't want to handle the error, 
	// you can use _ to ignore the error otherwise writing 'input, err := reader.ReadString('\n')' will also work ..
	// but you have to handle the error otherwise not handling err would give an error.
	
	fmt.Println("Thanks for rating, ", input) // Print.ln injects a new line '\n' after the output.
}
