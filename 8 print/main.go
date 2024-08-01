package main

import "fmt"

func main() {
	age := 25
	name := "Ram"
	height := 77.9988999
  
	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	fmt.Println("Hello world")

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)


	fmt.Printf("name is %s\n", name)
	fmt.Printf("Age is %d\n", age)
	fmt.Print("Name: %s Age: %d, Height: %.2f\n", name, age, height)

}