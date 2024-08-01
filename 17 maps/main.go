package main


import (
	"fmt"
)

func main() {

	// name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Prince"] = 100
	studentGrades["Alice"] = 89
	studentGrades["Bob"] = 90
	studentGrades["Piyush"] = 23

	fmt.Println("Marks of Bob: ", studentGrades["Bob"])
	studentGrades["Bob"] = 100
	fmt.Println("New Marks of Bob: ", studentGrades["Bob"])

	//Delete
	delete(studentGrades, "Bob")
	fmt.Println("Marks of Bob delete: ", studentGrades["Bob"])

	//Changing if a key exists
	grads, exists := studentGrades["David"]
	fmt.Println("Grads of David: ", grads)
	fmt.Println("David exists: ", exists)


  //Changing if a key exists
	Grads, Exists := studentGrades["Prince"]
	fmt.Println("Grads of Prince: ", Grads)
	fmt.Println("Prince exists: ", Exists)

	//loop
	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}


	//when map creation that time initization
	person := map[string]int{
		"Alice": 90,
		"Bob": 85,
		"Charlie": 95,
	}

	for index, value := range person {
		fmt.Printf("----------------Key is %s and marks is %d\n", index, value)
	}


}