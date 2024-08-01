package main

import (
	"fmt"
)


func main() {
	fmt.Println("we are learning Array in Golang")
/*
	//creating Array
	var name[5]string

	// 0 1 2 3 4
	name[0] = "prince"
	name[1] = "Aakash"
	name[2] = "Bharat"


	fmt.Println("Names of Person is: ", name)

	//creating Array another way
	var numbers = [5] int{1, 2, 3, 4, 5}
	fmt.Print("Number is: ", numbers)

	//length
	fmt.Println("Length of Numbers array is: ", len(numbers))

	fmt.Println("Value of name at 2 index is: ", name[2])
*/

	//""
	var name[5]string
	name[0] = "Prince"
	name[1] = "aakash"

	fmt.Println("name is: ", name)
	fmt.Printf("name Array is %q\n", name)
}