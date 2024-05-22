package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	// Declaring a map
	// Map is similar to a dictionary in Python
	languages := make(map[string]string) // This will create a map with key as string and value as string

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("Languages:", languages) // This will print map[JS:JavaScript PY:Python RB:Ruby]. The order may vary as maps are unordered
	// They are not comma separated as in Python

	// Accessing a value from a map
	fmt.Println("JS is short for", languages["JS"]) // This will print JavaScript

	// Deleting a key-value pair from a map
	delete(languages, "RB") // This will delete the key-value pair with key RB
	fmt.Println("Languages:", languages) // This will print map[JS:JavaScript PY:Python]

	// Checking if a key exists in a map
	_, ok := languages["PY"] // This will return true if the key PY exists in the map
	fmt.Println("PY exists in the map:", ok) // This will print true

	_, ok = languages["RB"] // This will return false if the key RB does not exist in the map
	fmt.Println("RB exists in the map:", ok) // This will print false

	// Iterating over a map
	for key, value := range languages {
		fmt.Println(key, "is short for", value)
	}

	// Declaring and initializing a map
	// You can declare and initialize a map in the following way

	// 1. var myMap = map[keyType]valueType{key1: value1, key2: value2}

	// 2. myMap := map[keyType]valueType{key1: value1, key2: value2}

	// 3. myMap := make(map[keyType]valueType)
	// myMap[key1] = value1
	// myMap[key2] = value2
	


}
