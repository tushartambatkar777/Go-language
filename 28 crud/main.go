package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Todo struct {
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/tod")
	if err != nil {
		fmt.Println("Error getting: ", err)
		return 
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response: ", res.Status)
		return
	}

	/*
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading: ", err)
		return
	}

	fmt.Println("Data: ", string(data))
	*/

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding: ", err)
		return
	}
	fmt.println("Todo: ", todo)



	fmt.Println("Title response: ", todo.Title)
	fmt.Println("Completed response: ", todoCompleted)
}

func performPostRequest() {
	todo := Todo {
		UserID: 23,
		Title: "Prince Kumar",
		Completed: true,
	}

	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return 
	}

	//convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonReader)

	myURL := "https://jsonplaceholder.typeicode.com/todos"

	//send POST request
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))

}

func performUpdateRequest() {
	todo := Todo {
		UserID: 23456,
		Title: "Tushar Kumar Golang World",
		Completed: false, 
	}
	
	// Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return 
	}

	// Convert json data to string
	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)


	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT Request
	http.NewRequest(http.MethodPUT, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT Request: ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return 
	}
	defer res.Body.Close()


	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response: ", string(data))
	fmt.Println("Response status: ", res.Status)

}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"


	//create DELETE Request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request: ", err)
		return
	}
	
	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return 
	}
	defer res.Body.Close()


	fmt.Println("Response status: ", res.Status)



}

func main() {
	fmt.Println("Learning CRUD")
	// performGetRequest()
	// performPostRequest()
	performUpdateRequest()
}