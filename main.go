package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Learn Go language")

	var name string = "Tushar"
	var version = 76
	fmt.Println(name)
	fmt.Println(version)

	var money int = 72000
	var currency = 3500
	fmt.Println(money)
	fmt.Println("Currency: ", currency)

	var dimension float64 = 99.99
	fmt.Println(dimension)
	
	var decided bool = false
	decided = true
	fmt.Println(decided)

	var person = 23
	fmt.Println(person)

	const pi = 3.14
	fmt.Println(pi)

	// another way to decliare variable
	person := "Kunal Nikam"
	fmt.Println(person)

	

	var Public = "data is important"
	var private = "data is private"
	fmt.Println(Public)
	fmt.Println(private)

}