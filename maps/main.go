package main

import "fmt"

func main() {
	// Map like in other languages is used to hold key-value pairs
	// Syntax : map[key type] value type{key : value}
	var country_and_city = map[string]string{"USA": "New York", "India": "Delhi", "Japan": "Tokyo", "China": "Shanghai"}
	fmt.Println(country_and_city)
	fmt.Println(country_and_city["USA"]) // get the values using the key

	num_and_alphabet := map[uint8]string{1: "a", 2: "b", 3: "c", 4: "d"}
	fmt.Println(num_and_alphabet)
	fmt.Println(num_and_alphabet[1])

	// Using make(map[key type] value type)
	sweet_fruits := make(map[string]bool)

	sweet_fruits["apple"] = true
	sweet_fruits["banana"] = true
	sweet_fruits["lime"] = false
	sweet_fruits["bitter kola"] = false
	sweet_fruits["dates"] = true
	sweet_fruits["mango"] = true
	fmt.Println(sweet_fruits)
	fmt.Println(sweet_fruits["apple"])
	fmt.Println(sweet_fruits["lime"])

	// In cases where it is not sure if the key-value exists. if the check is not implemented then Go will return a
	// default value for the type e.g "" for strings, 0 for all numerics, false for bool.
	fruit, is_present := sweet_fruits["mango"]
	if is_present {
		fmt.Println(fruit)
	} else {
		fmt.Println("Key not found")
	}
	fmt.Println()

	// Using for to print out the data
	// only keys
	for key := range sweet_fruits {
		fmt.Println(key)
	}
	fmt.Println()

	// only values
	for _, value := range sweet_fruits {
		fmt.Println(value)
	}
	fmt.Println()

	// both keys and values
	user_age := map[string]uint8{"Sam": 24, "Cornelius": 27, "Joseph": 24}
	for name, age := range user_age {
		fmt.Printf("%s is %d years old\n", name, age)
	}

}
