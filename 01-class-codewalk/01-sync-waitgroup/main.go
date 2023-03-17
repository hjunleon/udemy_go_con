package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup   // this is a semaphore essenntially
	wg.Add(1)

	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()
	fmt.Printf("the value of data is %v\n", data)
	// if data == 0 {
	// }
}
