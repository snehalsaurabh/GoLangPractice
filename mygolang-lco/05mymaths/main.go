package main

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	//"time"
)

func main() {
	fmt.Println("Welcome to maths in Golang")

	var mynumberOne int = 2
	var mynumberTwo float64 = 4.5
	fmt.Println("The sum of", mynumberOne, "and", mynumberTwo, "is", float64(mynumberOne) + (mynumberTwo)) // Type Conversion
	
	fmt.Println("Random - Refer Math/rand and Crypto/rand on pkg.go.dev") 
	// Using math/rand
	//" rand.Seed(time.Now().UnixNano()) "// Seed the random number generator - We use the current time in nanoseconds as the seed
	//Rand.Seed is depracated, use rand.NewSource(time.Now().UnixNano()) instead
	//" fmt.Println(rand.Intn(5)) " // This will generate a random number between 0 and 5


	// Using crypto/rand
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5)) // You have to handle the error here
	fmt.Println(myRandomNumber)
}