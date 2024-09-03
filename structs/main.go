package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newName string) {
	(*pointerToPerson).firstName = newName
}

func main() {
	// alex := person{"Alex", "Anderson"}  //a way to create a struct, not safe because use the order of props and values
	// alex := person{firstName: "Alex", lastName: "Anderson"}  //another way to create a struct

	// var alex person //another way to create an empty struct and add values after
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@email.com",
			zipCode: 97209,
		},
	}

	jim.updateFirstName("Rian")
	jim.print()

	name := "Bill"

	fmt.Println(*&name)
}
