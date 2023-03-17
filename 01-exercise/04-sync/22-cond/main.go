package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.RWMutex{} //https://stackoverflow.com/questions/48127289/rwmutex-in-conditional-variable
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		defer c.L.Unlock()
		// mu.RLock() // seems like will fk up
		// defer mu.RUnlock()
		for len(sharedRsc) == 0 {
			// time.Sleep(1 * time.Millisecond)
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		// c.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		defer c.L.Unlock()
		// mu.RLock()
		// defer mu.RUnlock()
		//TODO: suspend goroutine until sharedRsc is populated.
		// c.L.Lock()
		for len(sharedRsc) == 0 {
			// time.Sleep(1 * time.Millisecond)
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		// c.L.Unlock()
	}()

	// writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	c.Broadcast()
	c.L.Unlock()
	wg.Wait()
}
