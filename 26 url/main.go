package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "https://example.com:8080/path/to/resourse?key1=value1&key2=value2"
	fmt.Println("Type of URL: %T\n", myURL)

	parsedURL, err := url.Parse(myURL)
	if(err != nil) {
		fmt.Println("Can't parse URL ", err)
		return
	}
	fmt.Printf("Type of URL: %T\n", parsedURL)


	fmt.Println("Schema: ", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Println("RawQuery: ", parsedURL.RawQuery)


	//Modifying URL componenets
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=iamtushar"

	//Constructing a URL string from a URL object
	newURL := parsedURL.String()

	fmt.Println("new URL: ", newURL)

}