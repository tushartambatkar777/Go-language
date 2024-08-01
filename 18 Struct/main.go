package main

import (
	"fmt"
)

// int string bool Person Contact Address
type Person struct {
	FirstName string
	LastName string
	Age int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var tushar Person
	// fmt.Println("Tushar Person: ", tushar)
	tushar.FirstName = "Tushar"
	tushar.LastName = "Tambatkar"
	tushar.Age = 24
	// fmt.Println("Tushar Person: ", tushar)


	//2nd method
	person1 := Person{
		FirstName: "Vikram",
		LastName: "Bhatra",
		Age: 27,
	}
	fmt.Println("Person 1: ", person1)


	//new keyword
	var person2 = new(Person)
	person2.FirstName = "Simran"
	person2.LastName = "Thakur"
	person2.Age = 26

	// fmt.Println("Person 2: ", person.FirstName)
	// fmt.Println("Person 2: ", person.LastName)
	// fmt.Println("Person 2: ", person.Age)


	
	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Tushar",
		LastName: "Tambatkar",
		Age: 26,
	}
	employee1.Person_Contact.Email = "contact@gmail.com"
	employee1.Person_Contact.Phone = "8956256237"

	employee1.Person_Address = Address{
		House: 12,
		Area: "Tarwad",
		State: "Maharashtra",
	}

	fmt.Println("Employee 1: ", employee1)
}