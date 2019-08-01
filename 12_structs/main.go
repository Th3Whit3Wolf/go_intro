package main

import (
	"fmt"
	"strconv"
)

// Person Struct Definition
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int
	firstName, lastName, city, gender string
	age                               int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "David",
		lastName:  "Karrick",
		city:      "Feltwell",
		gender:    "Male",
		age:       27,
	}
	person2 := Person{
		firstName: "Nori", lastName: "Bissett", city: "Feltwell", gender: "Female", age: 21,
	}
	fmt.Println(person1, person2)
	fmt.Println(person1.firstName, person1.lastName, "loves", person2.firstName, person2.lastName)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

	fmt.Println(person2.greet())

	person2.hasBirthday()
	person2.getMarried(person1.lastName)
	fmt.Println(person2.greet())

}
