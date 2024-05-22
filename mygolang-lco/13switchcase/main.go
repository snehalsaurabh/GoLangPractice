package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in Golang")

	// Random number - New Syntax
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	random := rand.New(source)

	// Generate a random number between 1 and 6 (inclusive)
	diceNumber := random.Intn(6) + 1
	fmt.Println("Dice Number:", diceNumber)

	switch diceNumber {
	case 1: 
		fmt.Println("Dice Number is 1 & you can open") // You can see that break statement is not required in Go
	case 2:
		fmt.Println("Dice Number is 2") // It automatically breaks the case
	case 3:
		fmt.Println("Dice Number is 3") // If you want to fallthrough, you can use fallthrough keyword
	case 4:
		fmt.Println("Dice Number is 4") // It will execute the next case as well
		fallthrough
	case 5:
		fmt.Println("Dice Number is 5") // It will execute the next case as well
		fallthrough
	case 6:
		fmt.Println("Dice Number is 6 and roll the dice again") // It will not execute the next case here as there is no fallthrough
	default:
		fmt.Println("Invalid Dice Number")
	}

}
