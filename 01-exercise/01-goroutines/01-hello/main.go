package main

import (
	"fmt"
	"time"
	// "sync"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// var wg sync.WaitGroup   // this is a semaphore essenntially
	// wg.Add(3)
	// goroutine function call
	go fun("goroutine")

	// goroutine with anonymous function
	go func(msg string) {
		fun(msg)
		// defer wg.Done()
	}("anon")


	// goroutine with function value call
	fv := fun
	go fv("fv")


	// wait for goroutines to end


	fmt.Println("Waiting for all")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done..")
}
