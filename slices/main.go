package main

import (
	"fmt"
	"time"
)

func time_measurement(slice []int, n int) time.Duration {
	time_var := time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since((time_var))
}

func main() {
	// Slices are like array but with more flexibility and dynamic usage
	var my_slice = []int16{1, 2, 3, 4, 5} // No need to specify the size
	fmt.Println(my_slice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(my_slice), cap(my_slice))

	my_slice = append(my_slice, 6) // adds a new element to the slice
	fmt.Println(my_slice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(my_slice), cap(my_slice))

	my_slice = append(my_slice, 7, 8, 9, 10)
	fmt.Println(my_slice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(my_slice), cap(my_slice))

	// creating a slice from other slices using the spread operator
	var my_slice2 = []int16{11, 12, 13, 14, 15}
	fmt.Println(my_slice2)

	my_slice = append(my_slice, my_slice2...)
	fmt.Println(my_slice)
	fmt.Printf("The size of the slice is %d and the capacity is %d\n\n", len(my_slice), cap(my_slice))

	// Using make(type, size, capacity) for specifics. This is good if the size and capacity needed is known
	var my_slice3 = make([]int16, 5, 10)
	my_slice3[0] = 71
	my_slice3[1] = 72
	my_slice3[2] = 73
	my_slice3[3] = 74
	my_slice3[4] = 75
	fmt.Println(my_slice3)

	// Add the element from an array using a loop
	var my_slice4 []int16

	arr := [5]int16{1, 2, 3, 4, 5}
	for i := range len(arr) {
		my_slice4 = append(my_slice4, int16(arr[i]+2)) // adds 2 to each element before appending to the slice
	}
	fmt.Println(my_slice4)

	var my_slice5 []int16
	my_slice5 = append(my_slice5, arr[:]...) // [:] expands the array, ... appends its element to the slice
	fmt.Println(my_slice5)

	// Benchmark for specifying capacity vs not specifying capacity
	var x int = 1000000
	slice_a := []int{}                                                           // The compiler has to determine the capacity each time on its own
	slice_b := make([]int, 0, x)                                                 // capacity is preallocated, the compiler does'nt have to do that
	fmt.Printf("Time without preallocation: %s\n", time_measurement(slice_a, x)) // slower
	fmt.Printf("Time with preallocation: %s\n", time_measurement(slice_b, x))    // much faster during my tests

}
