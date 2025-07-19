package main

import "fmt"

func main() {
	const count int = 5
	i := 0

	// Without index
	for i < count {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println() // to add a new line

	for range count {
		fmt.Println("Hello")
	}

	// With index usage
	for i := 1; i <= count; i++ {
		fmt.Printf("%d. Hey \n", i)
	}
	fmt.Println()

	// Count down
	for i = 10; i >= 0; i-- {
		fmt.Println(i)
	}
}
