package main

import "fmt"

// An interface is a set of method signature, it defines what a type does, rather than how it does it. it has a close relationship with
// struct, and this combination allows Go to have the ability to mimic polymorphism in OOP even though it is not object-oriented.

type Animal interface {
	sound() string
	move()
}

type Dog struct {
	name string
}

type Ostrich struct {
	name string
}

// Methods
func (dog Dog) sound() string {
	return "barks"
}

func (dog Dog) move() {
	println("walk on 4 legs")
}

func (ostrich Ostrich) sound() string {
	return "chirps"
}

func (ostrich Ostrich) move() {
	println("Flightless, runs on 2 legs")
}

func check_animal(animal Animal) {
	if animal.sound() == "barks" {
		fmt.Println("this is a dog")
	} else {
		fmt.Println("this is an ostrich")
	}
}

func main() {
	dog_1 := Dog{name: "Edge"}
	ostrich_1 := Ostrich{name: "Colo"}

	// Calling the method individually
	fmt.Println(dog_1.sound())
	dog_1.move()
	fmt.Println()
	fmt.Println(ostrich_1.sound())
	ostrich_1.move()

	check_animal(dog_1)
	check_animal(ostrich_1)

	//
	animals := []Animal{dog_1, ostrich_1}
	for _, animal := range animals {
		fmt.Println(animal.sound())
		animal.move()
	}

}
