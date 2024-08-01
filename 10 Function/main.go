package main

import (
	"fmt"
)


func simpleFunction() {
	fmt.Println("I love YouTube")
}


func add(a, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("We are learning function in Golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println(ans)

	data := multiply(3, 4)
	fmt.Println("Multiplication of two numbers is :", data)
}