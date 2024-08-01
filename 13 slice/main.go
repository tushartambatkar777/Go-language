package main

import (
	"fmt"
)

func main() {
	/*
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 10, 12, 13)
	fmt.Println("Number : ", numbers)
	fmt.Printf("Number has data type : %T\n", numbers)
	fmt.Println("Lenght : ", len(numbers))

	name := []string{}
	fmt.Println("name : ", name)
	*/

	// Make func
	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 6)
	
	fmt.Println("Slices:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}