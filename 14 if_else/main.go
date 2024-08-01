package main

import (
	"fmt"
)

func main() {

	var num int;
	fmt.Scan(&num)
	
	/*
	if(num > 5) {
		fmt.Println("Number is greater than 5")
	} else {
		fmt.Println("Number is smaller than 5")
	}



		if(num > 10) {
			fmt.Println("Number is greater than 10")
		}else if(num > 5) {
			fmt.Println("Number is greater than 5 smaller than 10")
		}else{
			fmt.Println("Number is smaller than 5")
		}
		*/


		
		if(num < 5 && (num > 5 || num < 43)) {
			fmt.Println("Hey how are you!")
		}else{
			fmt.Println("Learn prigramming with Hello World")
		}
}