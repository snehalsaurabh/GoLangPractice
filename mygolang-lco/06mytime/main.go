package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in Golang")

	// Present Time
	presentTime := time.Now()
	fmt.Println("Present Time:", presentTime)

	// Formatting the time
	fmt.Println("Formatted Time:", presentTime.Format("2006-01-02 Monday 15:04:05"))
	// You have to remember the reference date: Mon, 15:04:05, 2006-01-02 as the reference date for formatting

	// Create a custom time
	createdDate := time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC) // Everthing is integer except the month and timezone
	fmt.Println("Created Date:", createdDate)
	// Formatting the custom time
	fmt.Println("Formatted Created Date:", createdDate.Format("2006-01-02 Monday 15:04:05"))
}
