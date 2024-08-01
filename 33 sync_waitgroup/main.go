package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worked %d started\n", i)
	fmt.Printf("Worker %d end\n", i)
}

func main() {

	var wg sync.WaitGroup
	//start 3 worker goroutines
	for i:=1; i<=3; i++ {
		wg.Add(1)  //Increment the WaitGroup counter
		go worker(i, &wg)
	}

	//Wait for all workers to finish
	wg.Wait()
	fmt.Println("Workers task complete")
}