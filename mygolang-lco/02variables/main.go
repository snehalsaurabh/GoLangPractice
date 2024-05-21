package main

import "fmt"

// jwtToken := 3000
// The walrus operation can only be used inside a function.
// Here, you will have to use a var keyword to declare the variable. 
// var jwtToken = 3000 is correct.
// jwtToken := 3000 // This will throw an error because the walrus operator can only be used inside a function.


// const
const LoginToken string = "asjdhjashdjkashd" // This is a constant variable. It can't be changed.
// The reason L in the LoginToken is in uppercase is because it is a public variable. This is equivalent to writing public in other languages.


func main() {
	//Strings
	var username string = "Snehal"
	// var is a keyword in Go that is used to declare variables.
	// If you declare a variable and don't use it, the Go compiler will throw an error.
	// The Go compiler is very strict about this.
	// Since we tried saving the file and did not use the username variable, the import "fmt" is not required and hence it was automatically removed.
	
	fmt.Println(username)
	// fp is the shortcut for fmt.Println
	// using the fmt.println function brought back the import "fmt" statement.

	//Now, we will try to use the fmt.Printf function instead of the println function
	fmt.Printf("Hello, I am %s.\n", username)
	fmt.Printf("Variable is of type %T\n", username)

	
	
	//Boolean
	var isLoggedIn bool = true // Boolean values are either true or false only. Can't put any other value.
	fmt.Println(isLoggedIn)
	fmt.Printf("Is the user logged in? %t\n", isLoggedIn)
	fmt.Printf("Variable is of type %T\n", isLoggedIn)

	
	
	//Integer - Small Value - uint8
	var smallVal uint8 = 255 // Only accepts value till 255. If you put 256, it will throw an error.
	// uint8 is an unsigned integer that can store values from 0 to 255.
	// You can simply use 'var smallVal = 255' as well. Go will automatically detect the type of the variable.
	// But it is always a good practice to declare the type of the variable.
	// This is because if you don't declare the type of the variable, Go will automatically assign the type of the variable based on the value assigned to it.
	// If you type 'var smallVal int = 255' also, it will handle most things but it is not a very good practice.
	
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T\n", smallVal)
	
	
	
	//Float - Small Value - float32
	var smallFloat float32 = 255.45544324234 // Only first 5 decimal places are shown but in float64, more (approx. 15) decimal places are shown.
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T\n", smallFloat)

	
	
	//Default Values
	var anotherVariable int //Default value is 0
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T\n", anotherVariable)

	
	
	//Implicit Type
	var website = "learncodeonline.in" // Go will automatically detect the type of the variable.
	fmt.Println(website)
	// Lexer decided the type of the variable based on the value assigned to it.
	// This is called implicit typing.
	// This is not a very good practice because it can lead to confusion.
	// Example of the confusion is given below:
	//"website = 4" // This will throw an error because the variable website is of type string and we are trying to assign an integer value to it.

	
	
	// No var keyword
	// This is the most common way of declaring variables in Go.
	// This is called the short declaration operator.
	numberOfUsers := 300 // This is the walrus operator - It is used to declare and assign a value to a variable.
	fmt.Println(numberOfUsers)


	// Calling the constant variable (public variable)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T\n", LoginToken)
}
