package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	// Values
	fmt.Println("Hi" + " Everyone") // concatenation
	fmt.Println("3.84 + 0.8 = ", 3.84+0.8)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// Variables
	// Variable declaration syntax = variable keyword | variable name | data type
	var age int8 = 26
	fmt.Println(age)
	fmt.Println("I am", age, "years old")

	age = 27
	fmt.Println("I will be", age, "years old next month")

	const pi float32 = 3.142
	fmt.Println(pi)

	var userName string = "Cornelius"
	fmt.Println("My name is", userName)

	var isGoFun bool = true
	fmt.Println("Is golang fun?", isGoFun)

	// Multiple initialization
	var a, b, c int = 1, 2, 3
	fmt.Println("a =", a, "b =", b, "c = ", c)

	var x float32 = 13.6779
	var y int8 = 9
	// var z float32 = x + y // Error
	var z float32 = x + float32(y) // Fixed using type conversion
	fmt.Println(z)

	var num1 int8 = 7
	var num2 int8 = 3
	fmt.Println(num1 / num2)                   // Rounds up the answer
	fmt.Println(num1 % num2)                   // Returns the remainder
	fmt.Println(float32(num1) / float32(num2)) // this will give the proper division answer expected

	var myRune rune = 'a'
	fmt.Println("character 'a' in ASCII is ", myRune) // outputs 97

	// Variable declaration without initialization
	var myVar string
	fmt.Println(myVar) // outputs empty string " "

	var num int
	fmt.Println(num) // outputs 0(same as all other numeric types uint8 - uint64, int8 - int64, floats)

	// Data types can be ignored, the compiler will infer it
	var animal = "Dog"
	fmt.Println(animal)

	animal = "Squirrel"
	fmt.Println(animal)

	var feline1, feline2 = "Cat", "Tiger"
	fmt.Println(feline1, feline2)

	// Even the var keyword can be ignored using this syntax
	// fruit string := "Apple" // Error!. Go does not allow specifying types for shorthand variable initialization
	fruit := "Apple"
	fmt.Println(fruit)

	name1, name2 := "Steve", "Wang"
	fmt.Println(name1, name2)

	// Constants
	const speedOfLight int32 = 299792458 // m/s
	fmt.Println("The speed of light is", speedOfLight)
	// speed_of_light = 2767990 // Error! can't reassign a constant

	result1 := add(4, 5) // In this case shorthand is not recommended since I'd want to know the return type of the function
	fmt.Println(result1)

	var result2 int = add(4, 6) // Better, the return type is indicated
	fmt.Println(result2)

}
