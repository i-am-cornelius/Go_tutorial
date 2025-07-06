package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	userName := "Cornelius"
	const message string = "Hi"
	fmt.Println(message, userName)

	userName = "Steve"
	greet(userName)

	var nums = [...]int16{78, 86, 71, 87, 94, 73, 82}
	for i := range nums {
		if nums[i]%2 != 0 {
			fmt.Println(nums[i], "is an odd number!")
		} else {
			fmt.Println(nums[i], "is an even number!")
		}
	}
}
