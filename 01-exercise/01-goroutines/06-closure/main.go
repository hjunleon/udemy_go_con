package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// go func() {
		// 	defer wg.Done()
		// 	fmt.Println(i) // references current value of i, which is 4 since main thread only enters waiting state after incrementingn i to 4
		// }()
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i) // paass by value 
	}
	wg.Wait()
}
