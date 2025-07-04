package main

import "fmt"

// Pass by value
func increment_1(arr [5]uint8) [5]uint8 {
	for i := range len(arr) {
		arr[i]++
	}
	fmt.Printf("Array address in increment_1() %p\n", &arr)
	return arr
}

// Pass by reference
func increment_2(arr *[5]uint8) [5]uint8 {
	for i := range len(arr) {
		arr[i]++
	}
	fmt.Printf("Address in increment_2() %p\n", arr)
	return *arr
}

func change_name(user_name *string) *string {
	*user_name = "Johnny"

	return user_name
}

func main() {
	var my_ptr *int32   // Declared but not initialized, pointer points to nil
	fmt.Println(my_ptr) // nil
	// fmt.Println(my_ptr) // Error! trying to read value address that doesn't exist, since my_ptr has not point to any address

	var my_var int32 = 10
	my_ptr = &my_var
	fmt.Println("The value is", my_var, "at address", my_ptr)

	// Dereferencing
	*my_ptr = 20
	fmt.Println("The value is", *my_ptr, "at address", my_ptr)

	// Initialize a pointer with new() and take memory space to hold value
	var user_name *string = new(string)
	fmt.Println("The value is", *user_name, "at address", user_name) // No value at the moment, but the pointer points to a legit address

	*user_name = "Cornelius" // writes a value to that address
	fmt.Println("The value is", *user_name, "at address", user_name)

	// Initialize with a variable address
	is_go_fun := true
	var is_go_fun_ptr *bool = &is_go_fun
	fmt.Println("The value is", is_go_fun, "at address", is_go_fun_ptr)

	// If a short hand syntax variable is used, then it is better to also use short hand for a pointer, to avoid type error
	a := 3.142
	a_ptr := &a
	fmt.Println("The value is", a, "at address", a_ptr)

	// Reassigning pointers
	var x float32 = 20.001
	var x_ptr *float32 = &x
	fmt.Println("The value is", x, "at address", x_ptr)

	var y float32 = 30.002
	x_ptr = &y
	fmt.Println("The value is", y, "at address", x_ptr)
	fmt.Println()

	// Function with pointers
	// Copying (pass by value) leading to different address
	nums := [5]uint8{1, 2, 3, 4, 5}
	fmt.Printf("%d Array address in main() %p\n", nums, &nums)
	fmt.Println(increment_1(nums)) // Incremented copies inside increment_1()
	fmt.Println(nums)              // The original elements remained unchanged
	fmt.Println()

	// pass by reference (no copies), same address
	fmt.Printf("%d Array address in main() %p\n", nums, &nums)
	fmt.Println(increment_2(&nums)) // Accepts the memory address of the array and modifies the original elements of the array
	fmt.Println(nums)               // The original elements show the change

	name := "Tom"
	fmt.Println(name)
	change_name(&name)
	fmt.Println(name)

	name_1 := "Ade"
	fmt.Println(name_1)
	change_name(&name_1)
	fmt.Println(name_1)

}
