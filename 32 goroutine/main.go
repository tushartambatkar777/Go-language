package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world")
	fmt.Println("Hello function ended successfully")
}

func sayHi() {
	fmt.Println("Hi Prince :)")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hi Prince Function ended")
}

func main() {
	fmt.Println("Learning goroutimes")
	go sayHello()
	sayHi()

	//Wait for a moment to allow the goroutime to finish
	time.Sleep(1000 * time.Millisecond)
}