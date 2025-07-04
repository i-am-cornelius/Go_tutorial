package main

import "fmt"

func main() {
	// Array structure in Go is similar to the C-style array
	// Declaration
	var nums [4]int8
	fmt.Println(nums) // filled with [0 0 0 0]

	nums[2] = 85
	fmt.Println(nums) // filled with [0 0 85 0]

	nums[0] = 81
	nums[1] = 82
	nums[3] = 83
	fmt.Println(nums) // filled with [81 82 85 83]

	fmt.Println(nums[1:3]) // prints elements at index 1 and 3 [82 85]

	// Memory addresses
	fmt.Println(&nums) // &[81 82 85 83]
	fmt.Println(&nums[0])
	fmt.Println(&nums[1])
	fmt.Println(&nums[2])
	fmt.Println(&nums[3])

	// Using loops print individually rather than as a collection of values
	for i := range nums {
		fmt.Printf("%d. %d \n", i, nums[i])
	}

	// Array initialization
	var fruits = [3]string{"apple", "orange", "banana"}
	fmt.Println(fruits)

	for i := range len(fruits) {
		fmt.Println(fruits[i])
	}

	animals := [4]string{"dog", "cat", "bird"} // an empty space is added for the last element not added
	fmt.Println(animals)

	//arr := [3]int{1, 2, 3, 4} // Error. Out of bounds
	arr := [...]int8{81, 72, 93, 84, 75, 66, 87, 78, 89} // Dynamic size.
	fmt.Println(arr)

	for index, element := range arr {
		fmt.Printf("%v. %v \n", index, element)
	}

}
