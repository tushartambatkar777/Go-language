package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}
func main() {
	fmt.Println("We are learning JSON")
	person := Person{Name: "Ram", Age: 18, IsAdult: true}
	//fmt.Println("Person Data is: ", person)

	// convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling ", err)
		return
	}
	fmt.Println("Person Data is: ", string(jsonData))

	
	// Decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error unmarshalling ", err)
		return
	}
	fmt.Println("person Data is after unmarshalling: ", personData)

}