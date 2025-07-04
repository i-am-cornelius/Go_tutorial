package main

import "fmt"

// Struct allows bundling together data that relate to each other in usage, essentially creating a new data type

type Person struct {
	first_name string
	last_name  string
	age        uint8
	is_male    bool
	career     CareerPath

	// properties can be added just by type, without a name
	int
	Religion
}

type CareerPath struct {
	person_career string
}

type Religion struct {
	religion string
}

// Methods(The syntax is different from normal functions because methods are bound to the receiver )
func (person Person) greet() {
	fmt.Println("Hello", person.first_name, person.last_name)
}

func (person Person) get_age() (string, uint8, string) {
	user_name := fmt.Sprintf("%s %s is", person.first_name, person.last_name)
	return user_name, person.age, "years old"
}

// more logic usage
func canVote(person Person, age uint8) string {
	if age >= 18 {
		return person.first_name + " can vote"
	}
	return person.first_name + " cannot vote"
}

func main() {
	person_1 := Person{"John", "Doe", 100, true, CareerPath{"used in software examples"}, 1, Religion{"Christian"}}
	fmt.Println(person_1)
	fmt.Println(person_1.first_name)
	fmt.Println(person_1.age)

	person_1.last_name = "Thompson" // modifiable from the outside
	fmt.Println(person_1)

	// A struct can be created on the go for a one-time use, so no reusability
	person_2 := struct {
		first_name string
		last_name  string
		age        uint8
		is_male    bool
		career     CareerPath
	}{first_name: "Jane", last_name: "Doe", age: 20, is_male: false, career: CareerPath{"Also used in software examples"}}
	fmt.Println(person_2)
	fmt.Println(person_2.first_name)

	person_1.greet()
	fmt.Println(person_1.get_age())

	if person_2.is_male {
		fmt.Println(person_2.first_name, "is a male")
	} else {
		fmt.Println(person_2.first_name, "is a female")
	}

	person_3 := Person{"Sam", "Ade", 15, true, CareerPath{"Chef"}, 2, Religion{"Atheist"}}
	fmt.Println(canVote(person_3, person_3.age))
	fmt.Println(canVote(person_1, person_1.age))
	// fmt.Println(canVote(person_2, person_2.age)) // Error! since person_2 is not a Person struct. but the value person_2.age can be used

}
