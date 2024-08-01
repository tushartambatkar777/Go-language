package main

import (
	"fmt"
	
)

func modifyValueByReference(num *int) {
	*num = *num + 20
}


func main() {
	/*
	var num int
	num = 2

	var ptr *int
	ptr = &num

	fmt.Println("The num has value: ", num)
	fmt.Println("Pointer contains:", ptr)


//Alternate for above program
  num := 2
	ptr := &num

	fmt.Println("Pointer contains: ", ptr)
	fmt.Println("Data contains through Pointer:", *ptr)

*/

	var pointer *int
	if pointer == nil { // default pointer == nil
		fmt.Println("Pointer is not assigned")
	}



	value := 5
	modifyValueByReference(&value)  
  fmt.Println("Value contains: ", value)
}