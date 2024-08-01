package main

import (
	"fmt"
	
)

func main() {
/*
	//Basic for loop
	for i:=0; i<4; i++ {
		fmt.Println("Numbers is : ", i)
	}

	//Infinte loop
	counter := 0
	for {
		fmt.Println("Infinite loop")
		counter++
		if(counter == 3) {
			break;
		}
	}
*/

//Range keyword
	numbers := []int{11, 42, 83, 14, 75}
	for index, value := range numbers {
		fmt.Printf("Index of Numbers is %d, and value is %d\n", index, value)
	}

	data := "Hello world"
	for index, value := range data {
		fmt.Printf("Index of Data %d, and value is %c\n", index, value)
	}
}