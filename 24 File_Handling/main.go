package main

import (
	"fmt"
	// "io"
	"os"
)

func main() {
/*
	file, err := os.Create("Example.text")
	if(err != nil) {
		fmt.Println("Enter while creting file: ", err)
		return 
	}
	defer file.Close()

	content := "Hello world by Tushar"
	_, errors := io.WriteString(file, content+"\n")
	if(errors != nil) {
		fmt.Println("Error while writing file: ", errors)
		return 
	}

	fmt.Println("Sucessfully created file")
*/



/*
	file, err := os.Open("Example.text")
	if(err != nil) {
		fmt.Println("Enter while creting file: ", err)
		return 
	}
	defer file.Close()

	//Create a buffer to read the file content
	buffer := make([]byte, 1024)

	//Read the file content into the buffer
	for{
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if (err != nil) {
			fmt.Println("Error while reading file", err)
			return 
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))
 	}
*/



		//Read the entire file into a byte slice
		content, err := os.ReadFile("Example.text")
		if(err != nil) {
			fmt.Println("Error while reading file ", err)
			return
		}
		fmt.Println(string(content))


} 