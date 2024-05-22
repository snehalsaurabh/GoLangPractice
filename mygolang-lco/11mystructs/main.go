package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in Golang; no classes; no super or parent

	// Declaring a struct
	Snehal := User{"Snehal", "snehalsaurabh.code@gmail.com", true, 20}
	fmt.Println(Snehal) // This will print {Snehal snehalsaurabh.code@gmail.com true 20}
	fmt.Printf("%+v\n", Snehal) // This will print {Name:Snehal Email:snehalsaurabh.code@gmail.com Status:true Age:20} 
	// +v will print the field names as well

	// Accessing a field from a struct
	fmt.Println("Name:", Snehal.Name) // This will print Name: Snehal
	fmt.Printf("Email: %v and Age: %v\n", Snehal.Email, Snehal.Age) // This will print Email: snehalsaurabh.code@gmail.com and Age: 20
	// %v is used to print the value of the variable
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
