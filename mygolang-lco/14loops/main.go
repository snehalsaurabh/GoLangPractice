package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"} // Slice of strings

	fmt.Println(days)

	// For Loop - Type 1
	for d:=0; d<len(days); d++ { //++d is not allowed in Go
		fmt.Println(days[d])
	}
	
	// For Loop - Type 2 - Better way to write for loop
	for i := range days {
		fmt.Println(days[i])
	}

	// For Loop - Type 3 - Using range and index/value
	for _, day := range days { // _ is used to ignore the index here or you can type in something instead of _ to use the index
		fmt.Println(day) // day is the value
	}



	// This is similar to while loop in other languages - while is not available in Go as a dedicated keyword
	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			break // This will break the loop
		} else if rougueValue == 3 {
			rougueValue++
			continue // This will skip the current iteration and move to the next iteration
		} else if rougueValue == 7 {
			rougueValue++
			goto lco // This will jump to the label lco
		}
		fmt.Println(rougueValue)
		rougueValue++
	}

	// Go does not have do-while loop


	
	// Go-To Statement
	// This is not recommended to use in Go
	// This is just for demonstration
	lco:
		fmt.Println("Learning Code Online")
}